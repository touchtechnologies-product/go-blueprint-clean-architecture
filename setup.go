package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app"
	"log"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	compRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
	staffRepo "github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/staff"
	//"io"
	//"log"
	//
	//"github.com/opentracing/opentracing-go"
	//"github.com/sirupsen/logrus"
	//goxid "github.com/touchtechnologies-product/xid"
	//"github.com/uber/jaeger-lib/metrics"
	//
	//"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app"
	//"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"io"
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
	company := companyService.New(cRepo, appConfig.Timezone)

	sRepo, err := staffRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	panicIfErr(err)
	staff := staffService.New(sRepo, appConfig.Timezone)

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