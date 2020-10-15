package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *companyin.UpdateInput) (err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := makeCompanyIDFilters(input.ID)

	company := &domain.Company{}
	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return util.RepoReadErr(err)
	}

	update := companyin.UpdateInputToCompanyDomain(input)
	company.Name = update.Name

	err = impl.repo.Update(ctx, filters, company)
	if err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
