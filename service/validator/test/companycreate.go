package test

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

func (suite *PackageTestSuite) TestCompanyCreateStructLevelValidationValid() {
	given := companyin.MakeTestCreateInput()

	suite.repo.On("Read", mock.Anything, []string{fmt.Sprintf("id:eq:%s", given.ID)}, &domain.Company{}).Once().Return(errors.New("error"))
	suite.repo.On("Read", mock.Anything, []string{fmt.Sprintf("name:eq:%s", given.Name)}, &domain.Company{}).Once().Return(errors.New("error"))

	err := suite.validator.Validate(given)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestCompanyCreateStructLevelValidationInvalid() {
	given := companyin.MakeTestCreateInput()
	given.ID = ""

	suite.repo.On("Read", mock.Anything, []string{fmt.Sprintf("id:eq:%s", given.ID)}, &domain.Company{}).Once().Return(errors.New("error"))
	suite.repo.On("Read", mock.Anything, []string{fmt.Sprintf("name:eq:%s", given.Name)}, &domain.Company{}).Once().Return(errors.New("error"))

	err := suite.validator.Validate(given)

	suite.Error(err)
}