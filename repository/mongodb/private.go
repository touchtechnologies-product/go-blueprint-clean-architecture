package mongodb

import (
	"reflect"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (repo *Repository) makeFilters(filters map[string]interface{}) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for k, v := range filters {
		switch v.(type) {
		case map[string]interface{}:
			bsonFilters[k] = repo.makeFilters(v.(map[string]interface{}))
			break
		default:
			bsonFilters[k] = v
			break
		}
	}
	return bsonFilters
}

func (repo *Repository) makePagingOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetLimit(int64(perPage))
	opts.SetSkip(int64(skip))
	return opts
}
