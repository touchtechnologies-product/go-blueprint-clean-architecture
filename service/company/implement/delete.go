package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *companyin.DeleteInput) (err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return util.RepoReadErr(err)
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
