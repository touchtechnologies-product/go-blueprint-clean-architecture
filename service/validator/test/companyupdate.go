package test

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

func (suite *PackageTestSuite) TestCompanyUpdateStructLevelValidationValid() {
	given := companyin.MakeTestUpdateInput()

	filters := []string{
		fmt.Sprintf("id:ne:%s", given.ID),
		fmt.Sprintf("name:eq:%s", given.Name),
	}
	suite.repo.On("Read", mock.Anything, filters, &domain.Company{}).Once().Return(errors.New("error"))

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCompanyUpdateStructLevelValidationInvalid() {
	given := companyin.MakeTestUpdateInput()

	filters := []string{
		fmt.Sprintf("id:ne:%s", given.ID),
		fmt.Sprintf("name:eq:%s", given.Name),
	}
	suite.repo.On("Read", mock.Anything, filters, &domain.Company{}).Once().Return(nil)

	err := suite.validator.Validate(given)

	suite.Error(err)
}

func (suite *PackageTestSuite) TestCompanyUpdateStructLevelValidationInvalidUnique() {
	given := companyin.MakeTestUpdateInput()
	given.ID = ""
	filters := []string{
		fmt.Sprintf("id:ne:%s", given.ID),
		fmt.Sprintf("name:eq:%s", given.Name),
	}
	suite.repo.On("Read", mock.Anything, filters, &domain.Company{}).Once().Return(errors.New("error"))

	err := suite.validator.Validate(given)

	suite.Error(err)
}
