package test

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"

	"github.com/stretchr/testify/suite"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util/mocks"
)

type PackageTestSuite struct {
	suite.Suite
	ctx         context.Context
	uuid        util.UUID
	validator   validator.Validator
	companyRepo *mocks.Repository
	staffRepo   *mocks.Repository
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	var err error
	suite.uuid, err = util.NewUUID()
	suite.NoError(err)
	suite.companyRepo = &mocks.Repository{}
	suite.staffRepo = &mocks.Repository{}
	suite.validator = validator.New(suite.companyRepo, suite.staffRepo)
}
