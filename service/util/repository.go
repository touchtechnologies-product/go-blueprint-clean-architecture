package util

import (
	"context"
)

//go:generate mockery -name=Repository
type Repository interface {
	List(ctx context.Context, opt *PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Read(ctx context.Context, filters map[string]interface{}, out interface{}) (err error)
	Update(ctx context.Context, filters map[string]interface{}, ent interface{}) (err error)
	Delete(ctx context.Context, filters map[string]interface{}) (err error)
	Count(ctx context.Context, filters map[string]interface{}) (total int, err error)

	Push(ctx context.Context, param *SetOpParam) (err error)
	Pop(ctx context.Context, param *SetOpParam) (err error)
	IsFirst(ctx context.Context, param *SetOpParam) (is bool, err error)
	CountArray(ctx context.Context, param *SetOpParam) (total int, err error)
	ClearArray(ctx context.Context, param *SetOpParam) (err error)
}