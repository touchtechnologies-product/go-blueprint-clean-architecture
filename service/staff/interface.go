package staff

import (
	"context"

	"content-service/service/util"
)

//go:generate mockery --name=Staff
type Service interface {
	List(ctx context.Context, opt *util.PageOption) (total int, items []*View, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (staff *View, err error)
	Update(ctx context.Context, ID string, input *CreateInput) (err error)
	Delete(ctx context.Context, ID string) (err error)
}

type Staff struct {
	validator util.Validator
	uuid      util.UUID
	repo      util.Repository
	timezone  string
}

type Wrapper struct {
	svc *Staff
}

func New(validator util.Validator, uuid util.UUID, staffRepo util.Repository, timezone string) Service {
	return &Wrapper{
		svc: &Staff{
			validator: validator,
			uuid:      uuid,
			repo:      staffRepo,
			timezone:  timezone,
		},
	}
}
