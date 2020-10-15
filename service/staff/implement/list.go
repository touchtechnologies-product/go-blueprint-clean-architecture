package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.StaffView, err error) {
	total, records, err := impl.repo.List(ctx, opt, &domain.Staff{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*out.StaffView, len(records))
	for i, record := range records {
		items[i] = out.StaffToView(record.(*domain.Staff))
	}

	return total, items, nil
}
