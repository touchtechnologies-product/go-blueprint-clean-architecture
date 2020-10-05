package staff

import (
	"context"
	"fmt"

	domain "blueprint/domain/staff"
	"blueprint/service/util"
)

func (impl *Staff) List(ctx context.Context, opt *util.PageOption) (total int, list []*View, err error) {
	total, items, err := impl.repo.List(ctx, opt, domain.Staff{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	list = make([]*View, len(items))
	for i, item := range items {
		list[i] = staffToView(item.(*domain.Staff))
	}

	return total, list, nil
}

func (impl *Staff) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	staff, err := impl.createInputToStaffDomain(input, impl.timezone)
	if err != nil {
		return "", util.ConvertInputToDomainErr(err)
	}

	ID, err = impl.repo.Create(ctx, staff)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return ID, nil
}

func (impl *Staff) Read(ctx context.Context, ID string) (view *View, err error) {
	staff := &domain.Staff{}
	filters := impl.makeIDFilters(ID)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	return staffToView(staff), nil
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

func (impl *Staff) makeIDFilters(ID string) (filters []string) {
	return []string{fmt.Sprintf("id:eq:%s", ID)}
}
