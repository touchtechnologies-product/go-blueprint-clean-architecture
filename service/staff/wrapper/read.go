package wrapper

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Read(ctx context.Context, input *staffin.ReadInput) (view *out.StaffView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Read")
	defer sp.Finish()

	sp.LogKV("staffID", input.StaffID)

	view, err = wrp.service.Read(ctx, input)

	sp.LogKV("view", view)
	sp.LogKV("err", err)

	return view, err
}
