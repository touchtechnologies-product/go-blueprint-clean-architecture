package staff

import (
	"net/http"

	"github.com/opentracing/opentracing-go"

	serviceStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"

	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/goerror/ginresp"
)

func (staff *Staff) UpdateStaff(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		c.Request.Context(),
		opentracing.GlobalTracer(),
		"handler.updateStaff",
	)
	defer span.Finish()

	ID := c.Param("ID")
	input := &serviceStaff.CreateInput{}
	if err := c.ShouldBind(input); err != nil {
		ginresp.RespValidateError(c, err)
		return
	}

	err := staff.service.Update(ctx, ID, input)
	if err != nil {
		ginresp.RespWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
