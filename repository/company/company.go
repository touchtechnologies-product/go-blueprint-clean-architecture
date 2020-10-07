package company

import (
	"context"
	"fmt"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/mongodb"
)

type Repository struct {
	*mongodb.Repository
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Repository, err error) {
	mongoDB, err := mongodb.New(ctx, uri, dbName, collName)
	if err != nil {
		return nil, err
	}
	return &Repository{mongoDB}, nil
}

func (repo *Repository) FindByName(ctx context.Context, name string) (company *domain.Company, err error) {
	filters := []string{fmt.Sprintf("name:eq:%s", name)}
	err = repo.Read(ctx, filters, company)
	return company, err
}
