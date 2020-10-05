package staff

import (
	"blueprint/app/view"
	"blueprint/service/util"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/touchtechnologies-product/goerror/ginresp"
)

// GetStaffsByCompany godoc
// @Summary List Staff By Company
// @Description Get all the existing Staff
// @Accept  json
// @Produce  json
// @Success 200 {array} staff.GetStaffsByCompanyOutput
// @Router /staffsByCompany [get]
func (staff *Staff) GetStaffsByCompany(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.getStaffsByCompany",
	)
	defer span.Finish()

	input := &util.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	total, items, err := staff.service.List(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	view.MakePaginatorResp(c, tot)
}
