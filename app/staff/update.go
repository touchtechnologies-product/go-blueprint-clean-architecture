package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

// Update godoc
// @Tags Staffs
// @Summary Update contents of a staff
// @Description Update staff with a given staff ID according to a given data
// @param staff_id path string true "Staff ID"
// @Param input body staffin.UpdateInput true "Input"
// @Param X-Authenticated-Userid header string true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.StaffView}
// @Success 204 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 409 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /staffs/{staff_id} [put]
func (ctrl *Controller) Update(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.staff.Update",
	)
	defer span.Finish()

	input := &staffin.UpdateInput{
		ID: c.Param("id"),
	}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	err := ctrl.service.Update(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
