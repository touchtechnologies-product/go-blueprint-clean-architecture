package staff

import (
	"fmt"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/touchtechnologies-product/goerror/ginresp"
)

func (staff *Staff) GetStaffsByCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.getStaffsByStaff",
	)
	defer span.Finish()

	input := &common.ListOption{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	list, err := staff.service.List(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &view.GetStaffsByCompanyOutput{
		ID: fmt.Sprintf("%d", list.Total),
		Name: fmt.Sprintf("%d", list.Total),
	})
}
