package staff

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (wrp *Wrapper) List(ctx context.Context, opt *util.PageOption) (total int, list []*View, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.List")
	defer sp.Finish()

	sp.LogKV("page", opt.Page)
	sp.LogKV("per_page", opt.PerPage)
	sp.LogKV("filters", opt.Filters)

	total, list, err = wrp.svc.List(ctx, opt)

	sp.LogKV("total", total)
	sp.LogKV("list", list)
	sp.LogKV("err", err)

	return total, list, err
}

func (wrp *Wrapper) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Create")
	defer sp.Finish()

	sp.LogKV("name", input.Name)
	sp.LogKV("companyID", input.CompanyID)
	sp.LogKV("tel", input.Tel)

	ID, err = wrp.svc.Create(ctx, input)

	sp.LogKV("ID", ID)
	sp.LogKV("err", err)

	return ID, err
}

func (wrp *Wrapper) Read(ctx context.Context, ID string) (view *View, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Read")
	defer sp.Finish()

	sp.LogKV("ID", ID)

	view, err = wrp.svc.Read(ctx, ID)

	sp.LogKV("view", view)
	sp.LogKV("err", err)

	return view, err
}

func (wrp *Wrapper) Update(ctx context.Context, ID string, input *CreateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Update")
	defer sp.Finish()

	sp.LogKV("name", input.Name)
	sp.LogKV("companyID", input.CompanyID)
	sp.LogKV("tel", input.Tel)

	err = wrp.svc.Update(ctx, ID, input)

	sp.LogKV("err", err)

	return err
}

func (wrp *Wrapper) Delete(ctx context.Context, ID string) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Delete")
	defer sp.Finish()

	sp.LogKV("ID", ID)

	err = wrp.svc.Delete(ctx, ID)

	sp.LogKV("err", err)

	return err
}
