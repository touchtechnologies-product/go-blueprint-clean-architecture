package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

// Read godoc
// @Tags Staffs
// @Summary Get a staff by staff ID
// @Description Response a staff data with a given staff ID
// @param staff_id path string true "Staff ID"
// @Produce  json
// @Success 200 {object} view.SuccessResp{data=out.StaffView}
// @Success 204 {object} view.SuccessResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /staffs/{staff_id} [get]
func (ctrl *Controller) Read(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.staff.Read",
	)
	defer span.Finish()

	input := &staffin.ReadInput{StaffID: c.Param("id")}

	staff, err := ctrl.service.Read(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeSuccessResp(c, http.StatusOK, staff)
}
