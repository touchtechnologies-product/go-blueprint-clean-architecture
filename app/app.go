package app

import (
	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"

	companyService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	staffService "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
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
	router.POST("/staff", app.staff.CreateStaff)
	router.PUT("/staff/:id", app.staff.UpdateStaff)
	router.GET("/staffsByCompany", app.staff.GetStaffsByCompany)

	router.POST("/company", app.company.CreateCompany)

	return app
}
