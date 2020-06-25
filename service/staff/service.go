package staff

import (
	"context"

	"github.com/touchtechnologies-product/goerror"
	goxid "github.com/touchtechnologies-product/xid"

	domainStaff "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/staff"
)

//go:generate mockery -name=Service
type Service interface {
	GetStaffsByCompany(ctx context.Context, input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error)
	CreateStaff(ctx context.Context, input *CreateStaffInput) (*domainStaff.Staff, goerror.Error)
	UpdateStaff(ctx context.Context, input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error)
}

type StaffService struct {
	xid               *goxid.ID
	staffRepository   staff.Repository
	companyRepository company.Repository
}

func New(xid *goxid.ID, r staff.Repository, c company.Repository) *StaffService {
	return &StaffService{
		xid:               xid,
		staffRepository:   r,
		companyRepository: c,
	}
}
