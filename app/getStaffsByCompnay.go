package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/touchtechnologies-product/goerror/ginresp"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/inout/staff"
	serviceStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

func (app *App) GetStaffsByCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.getStaffsByCompany",
	)
	defer span.Finish()

	input := &serviceStaff.GetStaffsByCompanyInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	staffs, err := app.staffService.GetStaffsByCompany(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &staff.GetStaffsByCompanyOutput{
		Staffs: staff.ToStaffsOutput(staffs),
	})
}
