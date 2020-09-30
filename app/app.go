package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/docs"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type App struct {
	staffService   staff.Service
	companyService company.Service
}

func New(staffService staff.Service, companyService company.Service) *App {
	return &App{
		staffService:   staffService,
		companyService: companyService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews - Youtube Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		apiRoutes.POST("/staff", app.CreateStaff)
		apiRoutes.PUT("/staff", app.UpdateStaff)
		apiRoutes.GET("/staffsByCompany", app.GetStaffsByCompany)
		apiRoutes.POST("/company", app.CreateCompany)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
