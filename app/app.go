package app

import (
	"github.com/gin-gonic/gin"

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
	router.POST("/staff", app.CreateStaff)
	router.PUT("/staff", app.UpdateStaff)
	router.GET("/staffsByCompany", app.GetStaffsByCompany)
	router.POST("/company", app.CreateCompany)

	return app
}
