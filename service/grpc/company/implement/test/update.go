package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (suite *PackageTestSuite) TestUpdateCompanyGPRCSuccess() {
	givenInput := companyin.MakeTestUpdateGRPCInput()
	companyDomain := &domain.Company{}
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, companyDomain).Once().Return(nil)

	givenCompany := companyin.MakeTestUpdateGRPCCompanyToCompanyDomain()
	suite.repo.On("Update", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	_, err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestUpdateCompanyGPRCError() {
	givenInput := companyin.MakeTestUpdateGRPCInput()
	companyDomain := &domain.Company{}
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)

	givenUpdateRepoError := errors.New("UPDATE REPO ERROR")
	expectedError := util.RepoUpdateErr(errors.New("UPDATE REPO ERROR"))

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, companyDomain).Once().Return(nil)

	givenCompany := companyin.MakeTestUpdateGRPCCompanyToCompanyDomain()
	suite.repo.On("Update", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(givenUpdateRepoError)
	_, err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.Equal(err, expectedError)
}

func (suite *PackageTestSuite) TestUpdateCompanyGPRCReadCompanyBeforeUpdateCompanyError() {
	givenInput := companyin.MakeTestUpdateGRPCInput()
	companyDomain := &domain.Company{}

	givenReadRepoError := errors.New("READ REPO ERROR")
	expectedError := util.RepoReadErr(errors.New("READ REPO ERROR"))

	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, companyDomain).Once().Return(givenReadRepoError)
	_, err := suite.service.Update(suite.ctx, givenInput)

	suite.Error(err)
	suite.Equal(err, expectedError)
}

func (suite *PackageTestSuite) TestUpdateCompanyGPRCValidateError() {
	givenInput := companyin.MakeTestUpdateGRPCInput()
	givenInput.Name = ""

	givenValidateError := errors.New("VALIDATION")
	expectedError := util.ValidationUpdateErr(errors.New("VALIDATION"))

	suite.validator.On("Validate", givenInput).Once().Return(givenValidateError)
	_, err := suite.service.Update(suite.ctx, givenInput)

	suite.Equal(err, expectedError)
}
