package wrapper

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Update(ctx context.Context, input *companyin.UpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.company.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)

	err = wrp.service.Update(ctx, input)

	sp.LogKV("err", err)

	return err
}
