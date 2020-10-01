package company

import (
	"context"
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/common"
)

//go:generate mockery --name=Company
type Service interface {
	List(ctx context.Context, opt *common.ListOption) (list *common.List, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (company *domain.Company, err error)
	Update(ctx context.Context, ID string, input *CreateInput) (err error)
	Delete(ctx context.Context, ID string) (err error)
}

type Company struct {
	companyRepo common.Repository
	timezone    string
}

func New(companyRepo common.Repository, timezone string) Service {
	return &Company{
		companyRepo: companyRepo,
		timezone:    timezone,
	}
}
