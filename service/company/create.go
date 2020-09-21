package company

import (
	"context"

	"github.com/touchtechnologies-product/goerror"

	domainCompany "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
)

type CreateCompanyInput struct {
	Name string `json:"name" binding:"required"`
}

func (service *CompanyService) CreateCompany(ctx context.Context, input *CreateCompanyInput) (*domainCompany.Company, goerror.Error) {
	newCompany := domainCompany.Create(service.xid.Gen(), input.Name)

	if err := service.companyRepository.Save(ctx, newCompany); err != nil {
		return nil, err
	}

	return newCompany, nil
}
