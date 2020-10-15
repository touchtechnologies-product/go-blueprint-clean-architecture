package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

// Create godoc
// @Tags Staffs
// @Summary Create a new staff
// @Description A newly created staff ID will be given in a Content-Location response header
// @Param input body staffin.CreateInput true "Input"
// @Param X-Authenticated-Userid header string true "User ID"
// @Accept json
// @Produce json
// @Success 201 {object} view.SuccessResp
// @Success 400 {object} view.ErrResp
// @Success 401 {object} view.ErrResp
// @Success 403 {object} view.ErrResp
// @Success 422 {object} view.ErrResp
// @Success 500 {object} view.ErrResp
// @Success 503 {object} view.ErrResp
// @Router /staffs [post]
func (ctrl *Controller) Create(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.staff.Create",
	)
	defer span.Finish()

	input := &staffin.CreateInput{}
	if err := c.ShouldBindJSON(input); err != nil {
		view.MakeErrResp(c, err)
		return
	}

	ID, err := ctrl.service.Create(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, ID)
}
