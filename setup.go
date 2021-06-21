package main

import (
	"context"
	validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"
	"io"
	"log"
	"time"

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

	cfg.ServiceName = appConfig.AppName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter = &jaegerConf.ReporterConfig{LogSpans: true}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
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
	staff := staffService.New(validator, sRepo, generateID)

	// Start GRPC server
	go app.NewGrpcServer(appConfig, generateID, cRepo, sRepo, validator)
	time.Sleep(1 * time.Second)

	return app.New(staff, company)
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
