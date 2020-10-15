package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.CompanyView, err error) {
	total, records, err := impl.repo.List(ctx, opt, &domain.Company{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*out.CompanyView, len(records))
	for i, record := range records {
		items[i] = out.CompanyToView(record.(*domain.Company))
	}

	return total, items, nil
}
