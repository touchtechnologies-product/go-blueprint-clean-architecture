package company

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"blueprint/app/view"
	service "blueprint/service/company"
)

func (company *Company) CreateCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createCompany",
	)
	defer span.Finish()

	input := &service.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	ID, err := company.service.Create(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, ID)
}
