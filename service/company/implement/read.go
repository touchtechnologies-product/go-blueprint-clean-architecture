package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *companyin.ReadInput) (view *out.CompanyView, err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.CompanyID)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return out.CompanyToView(company), nil
}
