package wrapper

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
)

func (wrp *wrapper) List(ctx context.Context, opt *domain.PageOption) (total int, list []*out.StaffView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.List")
	defer sp.Finish()

	sp.LogKV("page", opt.Page)
	sp.LogKV("per_page", opt.PerPage)
	sp.LogKV("filters", opt.Filters)

	total, list, err = wrp.service.List(ctx, opt)

	sp.LogKV("total", total)
	sp.LogKV("list", list)
	sp.LogKV("err", err)

	return total, list, err
}
