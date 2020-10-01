package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/touchtechnologies-product/goerror/ginresp"

	service "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

func (company *Company) CreateCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createCompany",
	)
	defer span.Finish()

	input := &service.CreateInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	ID, err := company.service.Create(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &view.CreateCompanyOutput{ID: ID})
}
