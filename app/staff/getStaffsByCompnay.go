package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

// GetStaffsByCompany godoc
// @Summary List Staff By Company
// @Description Get all the existing Staff
// @Accept  json
// @Produce  json
// @Success 200 {array} staff.View
// @Router /staffsByCompany} [get]
func (staff *Staff) GetStaffsByCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.getStaffsByCompany",
	)
	defer span.Finish()

	input := &util.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	total, items, err := staff.service.List(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}
