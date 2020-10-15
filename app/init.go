package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/docs"
	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type App struct {
	staff   *staff.Controller
	company *company.Controller
}

func New(staffService staffService.Service, companyService companyService.Service) *App {
	return &App{
		staff:   staff.New(staffService),
		company: company.New(companyService),
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
		apiRoutes.GET("/companies", app.company.Update)
		apiRoutes.POST("/companies", app.company.Create)
		apiRoutes.GET("/companies/:id", app.company.Read)
		apiRoutes.PUT("/companies/:id", app.company.Update)
		apiRoutes.DELETE("/companies/:id", app.company.Delete)
		
		apiRoutes.GET("/staffs", app.staff.Update)
		apiRoutes.POST("/staffs", app.staff.Create)
		apiRoutes.GET("/staffs/:id", app.staff.Read)
		apiRoutes.PUT("/staffs/:id", app.staff.Update)
		apiRoutes.DELETE("/staffs/:id", app.staff.Delete)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
