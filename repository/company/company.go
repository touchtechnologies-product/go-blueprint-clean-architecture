package company

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/base"
)

type Company struct {
	*base.MongoDB
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Company, err error) {
	mongoDB, err := base.New(ctx, uri, dbName, collName)
	if err != nil {
		return nil, err
	}
	return &Company{mongoDB}, nil
}