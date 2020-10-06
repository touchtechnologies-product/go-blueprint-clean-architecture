package company

import (
	"context"
	"fmt"

	"content-service/domain"
	"content-service/service/util"
)

func (impl *Company) List(ctx context.Context, opt *util.PageOption) (total int, list []*View, err error) {
	total, items, err := impl.repo.List(ctx, opt, domain.Company{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	list = make([]*View, len(items))
	for i, item := range items {
		list[i] = companyToView(item.(*domain.Company))
	}

	return total, list, nil
}

func (impl *Company) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	company := impl.createInputDomain(input)
	ID, err = impl.repo.Create(ctx, company)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return ID, nil
}

func (impl *Company) Read(ctx context.Context, ID string) (view *View, err error) {
	company := &domain.Company{}
	filters := impl.makeIDFilters(ID)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return companyToView(company), nil
}

func (impl *Company) makeIDFilters(ID string) (filters []string) {
	return []string{fmt.Sprintf("id:eq:%s", ID)}
}
