package base

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	coll   *mongo.Collection
	uri    string
	dbName string
}

func New(ctx context.Context, uri string, dbName string, collName string) (repo *MongoDB, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", uri, dbName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fullURI))
	if err != nil {
		return nil, err
	}

	mongoDB := &MongoDB{}
	mongoDB.uri = uri
	mongoDB.dbName = dbName
	mongoDB.client = client
	mongoDB.db = client.Database(dbName)
	mongoDB.coll = mongoDB.db.Collection(collName)

	return mongoDB, nil
}

func (db *MongoDB) List(ctx context.Context, opt *common.ListOption, itemType interface{}) (list *common.List, err error) {
	list = &common.List{}

	var filters bson.M
	var opts *options.FindOptions = nil
	if opt != nil {
		opts = db.makePagingOpts(opt.Page, opt.PerPage)
		if opt.Filters != nil && len(opt.Filters) > 0 {
			filters = db.makeFilters(opt.Filters)
		}
	}

	list.Total, err = db.count(ctx, filters)
	if err != nil {
		return nil, err
	}

	cursor, err := db.coll.Find(ctx, filters, opts)
	if err != nil {
		return nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var items []interface{}
	for cursor.Next(ctx) {
		item, err := db.clone(itemType)
		if err != nil {
			return nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	list.Items = &items
	if opt != nil {
		list.CurPage = opt.Page
		list.PerPage = opt.PerPage
		list.HasMore = (opt.Page * opt.PerPage) < list.Total
	}

	return list, nil
}

func (db *MongoDB) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (db *MongoDB) Create(ctx context.Context, ent interface{}) (ID string, err error) {
	res, err := db.coll.InsertOne(ctx, ent)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (db *MongoDB) Read(ctx context.Context, filters map[string]interface{}, out interface{}) (err error) {
	conds := db.makeFilters(filters)
	return db.coll.FindOne(ctx, conds).Decode(out)
}

func (db *MongoDB) Update(ctx context.Context, filters map[string]interface{}, ret interface{}) (err error) {
	conds := db.makeFilters(filters)
	_, err = db.coll.UpdateOne(ctx, conds, bson.M{"$set": ret})
	return err
}

func (db *MongoDB) Delete(ctx context.Context, filters map[string]interface{}) (err error) {
	conds := db.makeFilters(filters)
	_, err = db.coll.DeleteOne(ctx, conds)
	return err
}

func (db *MongoDB) Push(ctx context.Context, param *common.SetOpParam) (err error) {
	filters := db.makeFilters(map[string]interface{}{"id": param.ID})
	update := bson.M{
		"$addToSet": bson.M{
			param.SetFieldName: param.Item,
		},
	}
	_, err = db.coll.UpdateOne(ctx, filters, update)
	return err
}

func (db *MongoDB) Pop(ctx context.Context, param *common.SetOpParam) (err error) {
	filters := db.makeFilters(map[string]interface{}{"id": param.ID})
	update := bson.M{
		"$pop": bson.M{
			param.SetFieldName: -1,
		},
	}
	_, err = db.coll.UpdateOne(ctx, filters, update)
	return err
}

func (db *MongoDB) IsFirst(ctx context.Context, param *common.SetOpParam) (is bool, err error) {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"id": param.ID},
		},
		bson.M{
			"$project": bson.M{
				"id":    1,
				"first": bson.M{"$arrayElemAt": bson.A{fmt.Sprintf("$%s", param.SetFieldName), 0}},
			},
		},
	}

	cursor, err := db.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return false, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var items []struct {
		ID    primitive.ObjectID `bson:"_id,omitempty"`
		First string             `bson:"first"`
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return false, err
	}

	if len(items) < 1 {
		return false, err
	}

	return items[0].First == param.Item.(string), nil
}

func (db *MongoDB) CountArray(ctx context.Context, param *common.SetOpParam) (total int, err error) {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"id": param.ID},
		},
		bson.M{
			"$project": bson.M{
				"id":    1,
				"total": bson.M{"$size": fmt.Sprintf("$%s", param.SetFieldName)},
			},
		},
	}

	cursor, err := db.coll.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var items []struct {
		ID    string `bson:"id"`
		Total int    `bson:"total"`
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return 0, err
	}

	if len(items) < 1 {
		return 0, err
	}

	return items[0].Total, nil
}

func (db *MongoDB) ClearArray(ctx context.Context, param *common.SetOpParam) (err error) {
	filters := db.makeFilters(map[string]interface{}{"id": param.ID})
	_, err = db.coll.UpdateOne(ctx, filters, bson.M{"$set": bson.M{param.SetFieldName: param.Item}})
	return err
}

func (db *MongoDB) count(ctx context.Context, filter bson.M) (total int, err error) {
	cnt, err := db.coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}

	return int(cnt), nil
}

func (db *MongoDB) makeFilters(filters map[string]interface{}) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for k, v := range filters {
		switch v.(type) {
		case map[string]interface{}:
			bsonFilters[k] = db.makeFilters(v.(map[string]interface{}))
			break
		default:
			bsonFilters[k] = v
			break
		}
	}
	return bsonFilters
}

func (db *MongoDB) makePagingOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetLimit(int64(perPage))
	opts.SetSkip(int64(skip))
	return opts
}
