package wrapper

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Delete(ctx context.Context, input *staffin.DeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Delete")
	defer sp.Finish()

	sp.LogKV("staffID", input.ID)

	err = wrp.service.Delete(ctx, input)

	sp.LogKV("err", err)

	return err
}
