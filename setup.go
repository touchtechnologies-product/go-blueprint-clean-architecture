package main

import (
	"context"

	"github.com/uber/jaeger-client-go/rpcmetrics"

	"io"
	"log"

	companyWrapper "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/wrapper"
	staffWrapper "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/wrapper"
	validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"

	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	compRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
	staffRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/staff"
	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/implement"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/implement"
)

func setupJaeger(appConfig *config.Config) io.Closer {
	cfg, err := jaegerConf.FromEnv()
	panicIfErr(err)

	cfg.ServiceName = appConfig.AppName + "-" + appConfig.AppEnv
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{
		LogSpans:           true,
		LocalAgentHostPort: appConfig.JaegerAgentHost + ":" + appConfig.JaegerAgentPort,
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory
	jMetricsFactory = jMetricsFactory.Namespace(metrics.NSOptions{Name: appConfig.AppName + "-" + appConfig.AppEnv, Tags: nil})

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
		jaegerConf.Observer(rpcmetrics.NewObserver(jMetricsFactory, rpcmetrics.DefaultNameNormalizer)),
	)
	panicIfErr(err)
	opentracing.SetGlobalTracer(tracer)

	return closer
}
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	cRepo, err := compRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	panicIfErr(err)
	sRepo, err := staffRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	panicIfErr(err)

	validator := validatorService.New(cRepo, sRepo)
	generateID, err := util.NewUUID()
	panicIfErr(err)

	company := companyService.New(validator, cRepo, generateID)
	warpCompany := companyWrapper.WrapCompany(company)
	staff := staffService.New(validator, sRepo, generateID)
	wrapperStaff := staffWrapper.WrapperStaff(staff)
	return app.New(wrapperStaff, warpCompany)
}

func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
