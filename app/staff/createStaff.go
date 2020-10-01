package staff

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"net/http"

	"github.com/opentracing/opentracing-go"

	service "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"

	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/goerror/ginresp"
)

func (staff *Staff) CreateStaff(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.createStaff",
	)
	defer span.Finish()

	input := &service.CreateInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	ID, err := staff.service.Create(ctx, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, &view.CreateStaffOutput{ID: ID})
}
