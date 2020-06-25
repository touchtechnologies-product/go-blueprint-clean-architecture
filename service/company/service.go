package company

import (
	"context"

	"github.com/touchtechnologies-product/goerror"
	goxid "github.com/touchtechnologies-product/xid"

	domainCompany "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/company"
)

//go:generate mockery -name=Service
type Service interface {
	CreateCompany(ctx context.Context, input *CreateCompanyInput) (*domainCompany.Company, goerror.Error)
}

type CompanyService struct {
	companyRepository company.Repository
	xid               *goxid.ID
}

func New(xid *goxid.ID, c company.Repository) *CompanyService {
	return &CompanyService{
		companyRepository: c,
		xid:               xid,
	}
}
