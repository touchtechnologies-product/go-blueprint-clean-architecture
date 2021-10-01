package wrapper

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Read(ctx context.Context, input *companyin.ReadInput) (view *out.CompanyView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.company.Read")
	defer sp.Finish()

	sp.LogKV("id", input.CompanyID)

	view, err = wrp.service.Read(ctx, input)

	sp.LogKV("view", view)
	sp.LogKV("err", err)

	return view, err
}
