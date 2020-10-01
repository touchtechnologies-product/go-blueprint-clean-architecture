package staff

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/base"
)

type Staff struct {
	*base.MongoDB
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Staff, err error) {
	mongoDB, err := base.New(ctx, uri, dbName, collName)
	if err != nil {
		return nil, err
	}
	return &Staff{mongoDB}, nil
}