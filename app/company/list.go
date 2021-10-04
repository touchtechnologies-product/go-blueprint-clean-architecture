package company

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

// List godoc
// @Tags Companies
// @Summary Get a list of companies
// @Description Return a list of companies filtered by a given filters if any
// @param page query string false "A page number"
// @param per_page query string false "A total number of items per page"
// @param filters query []string false "Condition for company retrieval, ex. 'companyName:eq:some name'"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=[]out.CompanyView}
// @Success 204 {object} view.SuccessResp{data=[]out.CompanyView}
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies [get]
func (ctrl *Controller) List(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.company.List",
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

	span.LogKV("Input Handler", input)

	total, items, err := ctrl.service.List(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakePaginatorResp(c, total, items)
}
