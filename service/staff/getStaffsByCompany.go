package staff

import (
	"context"

	"github.com/touchtechnologies-product/goerror"

	domainStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
)

type GetStaffsByCompanyInput struct {
	CompanyId string `json:"companyId" form:"companyId" binding:"required"`
	Limit     int64  `json:"limit,default=20" form:"limit"`
	Offset    int64  `json:"offset" form:"offset"`
}

func (service *StaffService) GetStaffsByCompany(ctx context.Context, input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error) {
	staffs, err := service.staffRepository.GetStaffsByCompany(ctx, input.CompanyId, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}
