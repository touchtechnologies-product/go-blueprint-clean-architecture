package company

import (
	"blueprint/app/company"
	"blueprint/repository/mongodb"
	"context"
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

func (repo *Repository) FindByName(ctx context.Context, name string) (company *company.Company, err error) {
	filters := map[string]interface{}{"name": name}
	err = repo.Read(ctx, filters, company)
	return company, err
}
