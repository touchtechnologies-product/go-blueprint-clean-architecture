package staff

import (
	domain "blueprint/domain/staff"
	"blueprint/service/util"
	"context"
)

//go:generate mockery --name=Staff
type Service interface {
	List(ctx context.Context, opt *util.PageOption) (list *Paginator, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (staff *domain.Staff, err error)
	Update(ctx context.Context, ID string, input *CreateInput) (err error)
	Delete(ctx context.Context, ID string) (err error)
}

type Staff struct {
	validator util.Validator
	repo      util.Repository
	timezone  string
}

func New(validator util.Validator, staffRepo util.Repository, timezone string) Service {
	return &Staff{
		validator: validator,
		repo:      staffRepo,
		timezone:  timezone,
	}
}
