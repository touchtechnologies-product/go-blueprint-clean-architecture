package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	compImpl "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/implement"
	pbComp "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	staffImpl "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/implement"
	pbStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	validatorService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"
	"google.golang.org/grpc"
	"net"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/docs"
	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type App struct {
	staff   *staff.Controller
	company *company.Controller
}

const (
	NETWORK = "tcp"
)

func New(staffService staffService.Service, companyService companyService.Service) *App {
	return &App{
		staff:   staff.New(staffService),
		company: company.New(companyService),
	}
}

func NewGrpcServer(appConfig *config.Config, uuid util.UUID, compRepo util.Repository, staffRep util.Repository, validator validatorService.Validator) {
	companyServiceServer := compImpl.NewCompanyGrpcService(validator, compRepo, uuid)
	staffServiceServer := staffImpl.NewStaffGrpcService(validator, staffRep, uuid)

	lis, err := net.Listen(NETWORK, appConfig.GRPCAddress)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pbComp.RegisterCompanyGrpcServiceServer(grpcServer, companyServiceServer)
	pbStaff.RegisterStaffGrpcServiceServer(grpcServer, staffServiceServer)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	docs.SwaggerInfo.Title = "Touch Tech API"
	docs.SwaggerInfo.Description = "API Spec Demo."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		apiRoutes.GET("/companies", app.company.List)
		apiRoutes.POST("/companies", app.company.Create)
		apiRoutes.GET("/companies/:id", app.company.Read)
		apiRoutes.PUT("/companies/:id", app.company.Update)
		apiRoutes.DELETE("/companies/:id", app.company.Delete)

		apiRoutes.GET("/staff", app.staff.List)
		apiRoutes.POST("/staff", app.staff.Create)
		apiRoutes.GET("/staff/:id", app.staff.Read)
		apiRoutes.PUT("/staff/:id", app.staff.Update)
		apiRoutes.DELETE("/staff/:id", app.staff.Delete)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
