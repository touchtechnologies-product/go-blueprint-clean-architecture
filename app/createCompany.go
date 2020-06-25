package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/touchtechnologies-product/goerror/ginresp"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/inout/company"
	serviceCompany "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

func (app *App) CreateCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createCompany",
	)
	defer span.Finish()

	input := &company.CreateCompanyInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	newCompany, err := app.companyService.CreateCompany(ctx, &serviceCompany.CreateCompanyInput{Name: input.Name})
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &company.CreateCompanyOutput{
		Company: company.ToCompanyOutput(newCompany),
	})
}
