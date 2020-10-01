package base

import (
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/common"
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
	"time"
)

const collName = "RoomMock"

type MongoDBTestSuite struct {
	suite.Suite
	ctx context.Context
	db  *MongoDB
}

func (suite *MongoDBTestSuite) SetupSuite() {
	suite.ctx = context.Background()
}

func (suite *MongoDBTestSuite) SetupTest() {
	appConfig := config.Get()
	var err error
	suite.db, err = New(suite.ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, collName)
	suite.NoError(err)
	suite.NotNil(suite.db)
	suite.NoError(err)
}

func (suite *MongoDBTestSuite) TearDownTest() {
	_, _ = suite.db.coll.DeleteMany(suite.ctx, bson.M{})
}

func (suite *MongoDBTestSuite) TearDownSuite() {
	_ = suite.db.db.Drop(suite.ctx)
}

func TestMongoDBTestSuite(t *testing.T) {
	suite.Run(t, new(MongoDBTestSuite))
}

func (suite *MongoDBTestSuite) TestCreate() {
	ID, err := suite.db.Create(suite.ctx, domain.Create("cid", "cname"))
	suite.NoError(err)
	suite.NotEmpty(ID)
}

func (suite *MongoDBTestSuite) TestList() {
	company1 := domain.Create("cid", "cname")
	_, err := suite.db.Create(suite.ctx, company1)
	suite.NoError(err)
	time.Sleep(5*time.Second)
	company2 := domain.Create("cid", "cname")
	company2.Name = "cname2"
	_, err = suite.db.Create(suite.ctx, company2)
	suite.NoError(err)

	list := &common.List{}

	typeGuide := domain.Create("cid", "cname")
	list, err = suite.db.List(suite.ctx, common.MakeTestListOption(), typeGuide)

	suite.NoError(err)
	suite.Equal(2, list.Total)
}

func (suite *MongoDBTestSuite) TestRead() {
	company := domain.Create("cid", "cname")
	expect := domain.Create("cid", "cname")

	_, _ = suite.db.Create(suite.ctx, company)
	filters := map[string]interface{}{"id": company.Id}
	err := suite.db.Read(suite.ctx, filters, expect)

	suite.NoError(err)
	suite.Equal(company.Id, expect.Id)
}

func (suite *MongoDBTestSuite) TestUpdate() {
	company := domain.Create("cid", "cname")
	_, _ = suite.db.Create(suite.ctx, company)
	filters := map[string]interface{}{"id": company.Id}

	company.Name = "test update"
	err := suite.db.Update(suite.ctx, filters, company)
	suite.NoError(err)

	expect := domain.Create("cid", "cname")
	err = suite.db.Read(suite.ctx, filters, expect)

	suite.NoError(err)
	suite.Equal(company.Name, expect.Name)
}

func (suite *MongoDBTestSuite) TestDelete() {
	company := domain.Create("cid", "cname")
	filters := map[string]interface{}{"id": company.Id}
	_, _ = suite.db.Create(suite.ctx, company)
	err := suite.db.Delete(suite.ctx, filters)
	suite.NoError(err)
}

func (suite *MongoDBTestSuite) TestSearchList() {
	company1 := domain.Create("cid", "cname")
	company2 := domain.Create("cid", "cname")
	company2.Name = "test1"
	_, _ = suite.db.Create(suite.ctx, company1)
	_, _ = suite.db.Create(suite.ctx, company2)
	opt := common.MakeTestListOption()
	opt.Filters = map[string]interface{}{"name": map[string]interface{}{"$regex": "1"}}
	list, err := suite.db.List(suite.ctx, opt, domain.Create("cid", "cname"))
	suite.NoError(err)
	suite.Equal(1, list.Total)
}

func (suite *MongoDBTestSuite) TestPush() {
	company := domain.Create("cid", "cname")
	opt := &common.SetOpParam{
		ID:           "test",
		SetFieldName: "queue",
		Item:         "www",
	}

	_, _ = suite.db.Create(suite.ctx, company)
	err := suite.db.Push(suite.ctx, opt)
	suite.NoError(err)

	filters := map[string]interface{}{"id": company.Id}
	expect := domain.Create("cid", "cname")
	err = suite.db.Read(suite.ctx, filters, expect)

	suite.NoError(err)
	suite.Equal(company.Id, expect.Id)
}

func (suite *MongoDBTestSuite) TestNewWithIncorrectURI() {
	tmp := os.Getenv("MONGODB_URI")
	_ = os.Setenv("MONGODB_URI", "some-uri")
	appConfig := config.Get()
	var err error
	suite.db, err = New(suite.ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, collName)
	suite.Error(err)
	_ = os.Setenv("MONGODB_URI", tmp)
}

func (suite *MongoDBTestSuite) TestListCountErr() {
	_ = suite.db.client.Disconnect(suite.ctx)
	_, err := suite.db.List(suite.ctx, common.MakeTestListOption(), domain.Create("cid", "cname"))
	suite.Error(err)
}

func (suite *MongoDBTestSuite) TestCreateErr() {
	_, err := suite.db.Create(suite.ctx, nil)
	suite.Error(err)
}
