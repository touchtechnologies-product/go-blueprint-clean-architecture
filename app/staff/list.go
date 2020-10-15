package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

// List godoc
// @Tags Staffs
// @Summary Get a list of staffs
// @Description Return a list of staffs filtered by a given filters if any
// @param page query string false "A page number"
// @param per_page query string false "A total number of items per page"
// @param filters query []string false "Condition for staff retrieval, ex. 'staffName:eq:some name'"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]out.StaffView}
// @Success 204 {object} view.SuccessResp{data=[]out.StaffView}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /staffs [get]
func (ctrl *Controller) List(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.staff.List",
	)
	defer span.Finish()

	input := &domain.PageOption{}
	if err := c.ShouldBind(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	if len(input.Sorts) < 1 {
		input.Sorts = append(input.Sorts, "createdAt:desc")
	}

	total, items, err := ctrl.service.List(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}
