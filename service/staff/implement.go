package staff

import (
	"context"
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/common"
)


func (impl *Staff) List(ctx context.Context, opt *common.ListOption) (list *common.List, err error) {
	return impl.staffRepo.List(ctx, opt, domain.Staff{})
}

func (impl *Staff) Create(ctx context.Context, input *CreateInput) (ID string, err error) {
	staff, err := createInputToStaffDomain(input, impl.timezone)
	if err != nil {
		return "", err
	}
	return impl.staffRepo.Create(ctx, staff)
}

func (impl *Staff) Read(ctx context.Context, ID string) (staff *domain.Staff, err error) {
	staff = &domain.Staff{}
	filters := impl.makeIDFilters(ID)
	return staff, impl.staffRepo.Read(ctx, filters, staff)
}

func (impl *Staff) Update(ctx context.Context, ID string, input *CreateInput) (err error) {
	filters := impl.makeIDFilters(ID)
	return impl.staffRepo.Update(ctx, filters, input)
}

func (impl *Staff) Delete(ctx context.Context, ID string) (err error) {
	staff := &domain.Staff{}
	filters := impl.makeIDFilters(ID)
	err = impl.staffRepo.Read(ctx, filters, staff)
	if err != nil {
		return err
	}
	return impl.staffRepo.Delete(ctx, filters)
}

func (impl *Staff) makeIDFilters(ID string) (filters map[string]interface{}) {
	return map[string]interface{}{"id": ID}
}