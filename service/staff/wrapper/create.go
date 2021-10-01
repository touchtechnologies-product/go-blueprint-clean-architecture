package wrapper

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/opentracing/opentracing-go"
)

func (wrp *wrapper) Create(ctx context.Context, input *staffin.CreateInput) (ID string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Create")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)

	ID, err = wrp.service.Create(ctx, input)

	sp.LogKV("ID", ID)
	sp.LogKV("err", err)

	return ID, err
}
