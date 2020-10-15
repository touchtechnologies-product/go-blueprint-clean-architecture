package util

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

//go:generate mockery --name=Repository
type Repository interface {
	List(ctx context.Context, opt *domain.PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Update(ctx context.Context, filters []string, ent interface{}) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)

	Push(ctx context.Context, param *domain.SetOpParam) (err error)
	Pop(ctx context.Context, param *domain.SetOpParam) (err error)
	IsFirst(ctx context.Context, param *domain.SetOpParam) (is bool, err error)
	CountArray(ctx context.Context, param *domain.SetOpParam) (total int, err error)
	ClearArray(ctx context.Context, param *domain.SetOpParam) (err error)
}
