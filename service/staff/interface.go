package staff

import (
	"context"
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/common"
)

//go:generate mockery --name=Staff
type Service interface {
	List(ctx context.Context, opt *common.ListOption) (list *common.List, err error)
	Create(ctx context.Context, input *CreateInput) (ID string, err error)
	Read(ctx context.Context, ID string) (staff *domain.Staff, err error)
	Update(ctx context.Context, ID string, input *CreateInput) (err error)
	Delete(ctx context.Context, ID string) (err error)
}

type Staff struct {
	staffRepo common.Repository
	timezone  string
}

func New(staffRepo common.Repository, timezone string) Service {
	return &Staff{
		staffRepo: staffRepo,
		timezone:  timezone,
	}
}
