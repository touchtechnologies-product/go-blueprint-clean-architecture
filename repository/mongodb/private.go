package mongodb

import (
	"reflect"
	"strings"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (repo *Repository) makeFilters(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]
		value := slFilter[2]

		switch operations {
		case "ne":
			bsonFilters[key] = bson.M{"$ne": value}
			break
		case "like":
			bsonFilters[key] = bson.M{
				"$regex":   value,
				"$options": "i",
			}
			break
		case "eq":
			bsonFilters[key] = value
			break
		default:
			bsonFilters[key] = value
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
