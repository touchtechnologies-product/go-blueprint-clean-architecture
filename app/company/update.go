package company

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

// Update godoc
// @Tags Companies
// @Summary Update contents of a company
// @Description Update company with a given company ID according to a given data
// @param company_id path string true "Company ID"
// @Param input body companyin.UpdateInput true "Input"
// @Param X-Authenticated-Userid header string true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.CompanyView}
// @Success 204 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 409 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /companies/{company_id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.company.Update",
	)
	defer span.Finish()

	input := &companyin.UpdateInput{
		ID: c.Param("id"),
	}

	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}
	
	span.LogKV("Input Handler", input)

	err := ctrl.service.Update(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
