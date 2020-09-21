package staff

import (
	"context"

	"github.com/touchtechnologies-product/goerror"

	domainStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
)

type UpdateStaffInput struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Tel  string `json:"tel" binding:"required"`
}

func (service *StaffService) UpdateStaff(ctx context.Context, input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error) {
	staff, err := service.staffRepository.Get(ctx, input.Id)
	if err != nil {
		return nil, err
	}

	staff.Update(input.Name, input.Tel)

	if err := service.staffRepository.Save(ctx, staff); err != nil {
		return nil, err
	}

	return staff, nil
}
