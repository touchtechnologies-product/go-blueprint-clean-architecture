package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

// Delete godoc
// @Tags Staffs
// @Summary Soft delete a staff
// @Description Mark a staff with a given staff ID as deleted but keep its content
// @param staff_id path string true "Staff ID"
// @Param X-Authenticated-Userid header string true "User ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /staffs/{staff_id} [delete]
func (ctrl *Controller) Delete(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.staff.Delete",
	)
	defer span.Finish()

	input := &staffin.DeleteInput{
		ID:    c.Param("id"),
	}

	err := ctrl.service.Delete(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, nil)
}
