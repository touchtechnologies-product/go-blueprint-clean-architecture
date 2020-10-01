package company

import (
	"context"
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/common"
)


func (impl *Company) List(ctx context.Context, opt *common.ListOption) (list *common.List, err error) {
	return impl.companyRepo.List(ctx, opt, domain.Company{})
}

func (impl *Company) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	company, err := createInputToCompanyDomain(input, impl.timezone)
	if err != nil {
		return "", err
	}
	return impl.companyRepo.Create(ctx, company)
}

func (impl *Company) Read(ctx context.Context, ID string) (company *domain.Company, err error) {
	company = &domain.Company{}
	filters := impl.makeIDFilters(ID)
	return company, impl.companyRepo.Read(ctx, filters, company)
}

func (impl *Company) Update(ctx context.Context, ID string, input *CreateInput) (err error) {
	filters := impl.makeIDFilters(ID)
	return impl.companyRepo.Update(ctx, filters, input)
}

func (impl *Company) Delete(ctx context.Context, ID string) (err error) {
	company := &domain.Company{}
	filters := impl.makeIDFilters(ID)
	err = impl.companyRepo.Read(ctx, filters, company)
	if err != nil {
		return err
	}
	return impl.companyRepo.Delete(ctx, filters)
}

func (impl *Company) makeIDFilters(ID string) (filters map[string]interface{}) {
	return map[string]interface{}{"id": ID}
}