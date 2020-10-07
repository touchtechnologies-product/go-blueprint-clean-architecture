package mongodb

import (
	"blueprint/config"
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type TestStruct struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
	List  []string           `bson:"list"`
}

type MongoDBTestSuite struct {
	suite.Suite
	ctx  context.Context
	repo *Repository
}

func (suite *MongoDBTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	err := godotenv.Load()
	suite.NoError(err)
	conf := config.Get()

	suite.repo, err = New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBCompanyTableName)
	suite.NoError(err)
}

func (suite *MongoDBTestSuite) SetupTest() {
	err := godotenv.Load()
	suite.NoError(err)
	conf := config.Get()

	suite.repo, err = New(suite.ctx, conf.MongoDBEndpoint, conf.MongoDBName, conf.MongoDBCompanyTableName)
	suite.NoError(err)
}

func (suite *MongoDBTestSuite) TearDownTest() {
	_, _ = suite.repo.Coll.DeleteMany(suite.ctx, bson.M{})
}

func (suite *MongoDBTestSuite) TearDownSuite() {
	_ = suite.repo.DB.Drop(suite.ctx)
}

func TestMongoDBTestSuite(t *testing.T) {
	suite.Run(t, new(MongoDBTestSuite))
}

func (suite *MongoDBTestSuite) makeTestStruct(title string, list ...string) (test *TestStruct) {
	return &TestStruct{
		Title: title,
		List: list,
	}
}
