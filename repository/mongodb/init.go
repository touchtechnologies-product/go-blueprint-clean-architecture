package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
	uri    string
	dbName string
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *Repository, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", uri, dbName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	repo = &Repository{}
	repo.uri = uri
	repo.dbName = dbName
	repo.client = client
	repo.db = client.Database(dbName)
	repo.coll = repo.db.Collection(collName)

	return repo, nil
}
