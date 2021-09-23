package test

import (
	"context"
	"fmt"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/implement"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util/mocks"
	validatorMocks "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator/mocks"

	"github.com/stretchr/testify/suite"
)

// PackageTestSuite struct
type PackageTestSuite struct {
	suite.Suite
	ctx                 context.Context
	validator           *validatorMocks.Validator
	repo                *mocks.Repository
	uuid                *mocks.UUID
	service             company.Service
	makeDataTestCompany *domain.Company
}

func (suite *PackageTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.validator = &validatorMocks.Validator{}
	suite.repo = &mocks.Repository{}
	suite.service = implement.New(suite.validator, suite.repo, suite.uuid)
	suite.makeDataTestCompany = MakeTestCompany()
}

func (suite *PackageTestSuite) makeCompanyIDFilter(companyID string) (filters []string) {
	return []string{
		fmt.Sprintf("id:eq:%s", companyID),
	}
}
func MakeTestCompany() (company *domain.Company) {
	return &domain.Company{
		ID:   "test",
		Name: "test",
	}
}
