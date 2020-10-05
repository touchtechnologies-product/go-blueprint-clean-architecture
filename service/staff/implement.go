package staff

import (
	domain "blueprint/domain/staff"
	"blueprint/service/util"
	"context"
	"math"
)

func (impl *Staff) List(ctx context.Context, opt *util.PageOption) (list *Paginator, err error) {
	total, items, err := impl.repo.List(ctx, opt, domain.Staff{})
	if err != nil {
		return nil, util.RepoListErr(err)
	}

	list = &Paginator{}
	list.Items = make([]*domain.Staff, len(items))
	for i, item := range items {
		list.Items[i] = item.(*domain.Staff)
	}
	list.Total = total
	list.PerPage = opt.PerPage
	list.CurPage = opt.Page
	lastPage := int(math.Ceil(float64(total / opt.PerPage)))
	list.PageLeft = lastPage - opt.Page
	list.HasMore = opt.Page < lastPage

	return list, nil
}

func (impl *Staff) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	staff, err := createInputToStaffDomain(input, impl.timezone)
	if err != nil {
		return "", util.ConvertInputToDomainErr(err)
	}

	ID, err = impl.repo.Create(ctx, staff)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return ID, nil
}

func (impl *Staff) Read(ctx context.Context, ID string) (staff *domain.Staff, err error) {
	staff = &domain.Staff{}
	filters := impl.makeIDFilters(ID)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return staff, nil
}

func (impl *Staff) Update(ctx context.Context, ID string, input *CreateInput) (err error) {
	_, err = impl.Read(ctx, ID)
	if err != nil {
		return err
	}

	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	filters := impl.makeIDFilters(ID)
	err = impl.repo.Update(ctx, filters, input)
	if err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}

func (impl *Staff) Delete(ctx context.Context, ID string) (err error) {
	_, err = impl.Read(ctx, ID)
	if err != nil {
		return err
	}

	filters := impl.makeIDFilters(ID)
	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}

func (impl *Staff) makeIDFilters(ID string) (filters map[string]interface{}) {
	return map[string]interface{}{"id": ID}
}
