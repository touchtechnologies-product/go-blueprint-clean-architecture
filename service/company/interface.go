package company

import (
	"context"

	"blueprint/domain"
	"blueprint/service/util"
)

//go:generate mockery --name=Company
type Service interface {
	List(ctx context.Context, opt *util.PageOption) (total int, items []*View, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (company *View, err error)
}

//go:generate mockery --name=Repository
type Repository interface {
	util.Repository
	FindByName(ctx context.Context, name string) (company *domain.Company, err error)
}

type Company struct {
	validator util.Validator
	uuid      util.UUID
	repo      Repository
	timezone  string
}

type Wrapper struct {
	svc *Company
}

func New(validator util.Validator, uuid util.UUID, repo Repository, timezone string) Service {
	return &Wrapper{
		svc: &Company{
			validator: validator,
			uuid:      uuid,
			repo:      repo,
			timezone:  timezone,
		},
	}
}
