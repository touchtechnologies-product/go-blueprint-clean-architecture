package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *companyin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	company := companyin.CreateInputToCompanyDomain(input)

	_, err = impl.repo.Create(ctx, company)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return company.ID, nil
}
