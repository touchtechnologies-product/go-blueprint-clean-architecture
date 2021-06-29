package test

import (
	"context"
	"fmt"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/implement"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util/mocks"
	validatorMocks "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator/mocks"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite struct
type PackageTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *validatorMocks.Validator
	repo      *mocks.Repository
	uuid      *mocks.UUID
	service   company.Service
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.uuid = &mocks.UUID{}
	suite.repo = &mocks.Repository{}
	suite.service = implement.NewCompanyGrpcService(suite.validator, suite.repo, suite.uuid)
}

func (suite *PackageTestSuite) makeCompanyIDFilter(companyID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", companyID),
	}
}
