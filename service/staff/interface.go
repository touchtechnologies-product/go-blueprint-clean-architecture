package staff

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*out.StaffView, err error)
	Create(ctx context.Context, input *staffin.CreateInput) (ID string, err error)
	Read(ctx context.Context, input *staffin.ReadInput) (staff *out.StaffView, err error)
	Update(ctx context.Context, input *staffin.UpdateInput) (err error)
	Delete(ctx context.Context, input *staffin.DeleteInput) (err error)
}
