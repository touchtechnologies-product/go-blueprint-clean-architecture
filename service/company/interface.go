package company

import (
	domain "blueprint/domain/company"
	"blueprint/service/util"
	"context"
)

//go:generate mockery --name=Company
type Service interface {
	List(ctx context.Context, opt *util.PageOption) (list *Paginator, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (company *domain.Company, err error)
}

//go:generate mockery -name=CompanyRepository
type Repository interface {
	util.Repository
	FindByName(ctx context.Context, name string) (company *domain.Company, err error)
}

type Company struct {
	validator util.Validator
	repo      Repository
	timezone  string
}

func New(validator util.Validator, repo Repository, timezone string) Service {
	return &Company{
		validator: validator,
		repo:      repo,
		timezone:  timezone,
	}
}
