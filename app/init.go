package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"blueprint/app/company"
	"blueprint/app/staff"

	"blueprint/docs"
	companyService "blueprint/service/company"
	staffService "blueprint/service/staff"
)

type App struct {
	staff   *staff.Staff
	company *company.Company
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
		apiRoutes.POST("/staff", app.staff.CreateStaff)
		apiRoutes.PUT("/staff", app.staff.UpdateStaff)
		apiRoutes.GET("/staffsByCompany", app.staff.GetStaffsByCompany)
		apiRoutes.POST("/company", app.company.CreateCompany)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
