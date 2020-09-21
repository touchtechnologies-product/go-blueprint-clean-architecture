package app

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	staff2 "github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/inout/staff"
	serviceStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"

	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/goerror/ginresp"
)

func (app *App) CreateStaff(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createStaff",
	)
	defer span.Finish()

	input := &serviceStaff.CreateStaffInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staff, err := app.staffService.CreateStaff(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff2.CreateStaffOutput{
		Staff: staff2.ToStaffOutput(staff),
	})
}
