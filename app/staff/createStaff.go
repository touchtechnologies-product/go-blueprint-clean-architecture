package staff

import (
	"github.com/opentracing/opentracing-go"

	"blueprint/app/view"

	service "blueprint/service/staff"

	"github.com/gin-gonic/gin"
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
		view.MakeErrResp(c, err)
		return
	}

	ID, err := staff.service.Create(ctx, input)
	if err != nil {
		view.MakeErrResp(c, err)
		return
	}

	view.MakeCreatedResp(c, ID)
}
