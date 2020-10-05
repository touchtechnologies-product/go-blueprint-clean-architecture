package company

import (
	domain "blueprint/domain/company"
	"blueprint/service/util"
	"context"
	"math"
)

func (impl *Company) List(ctx context.Context, opt *util.PageOption) (list *Paginator, err error) {
	total, items, err := impl.repo.List(ctx, opt, domain.Company{})
	if err != nil {
		return nil, util.RepoListErr(err)
	}

	list = &Paginator{}
	list.Items = make([]*domain.Company, len(items))
	for i, item := range items {
		list.Items[i] = item.(*domain.Company)
	}
	list.Total = total
	list.PerPage = opt.PerPage
	list.CurPage = opt.Page
	list.LastPage = int(math.Ceil(float64(total / opt.PerPage)))
	list.HasMore = opt.Page < list.LastPage

	return list, nil
}

func (impl *Company) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	company := createInputDomain(input)
	ID, err = impl.repo.Create(ctx, company)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return ID, nil
}

func (impl *Company) Read(ctx context.Context, ID string) (company *domain.Company, err error) {
	company = &domain.Company{}
	filters := impl.makeIDFilters(ID)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return company, nil
}

func (impl *Company) makeIDFilters(ID string) (filters map[string]interface{}) {
	return map[string]interface{}{"id": ID}
}
