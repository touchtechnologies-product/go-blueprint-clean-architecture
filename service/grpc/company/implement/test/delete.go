package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (suite *PackageTestSuite) TestDeleteCompanyGPRCSuccess() {
	givenInput := companyin.MakeTestDeleteCompGRPCInput()
	givenCompany := &domain.Company{}
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)

	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	suite.repo.On("Delete", mock.Anything, givenCompanyIDFilter).Once().Return(nil)
	_, err := suite.service.Delete(suite.ctx, givenInput)

	suite.NoError(err)
}

func (suite *PackageTestSuite) TestDeleteCompanyGPRCReadError() {
	givenInput := companyin.MakeTestDeleteCompGRPCInput()
	givenCompany := &domain.Company{}

	givenReadRepoError := errors.New("READ REPO ERROR")
	expectedError := util.RepoReadErr(errors.New("READ REPO ERROR"))

	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)

	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(givenReadRepoError)
	_, err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.Equal(err, expectedError)
}

func (suite *PackageTestSuite) TestDeleteCompanyGPRCError() {
	givenInput := companyin.MakeTestDeleteCompGRPCInput()
	givenCompany := &domain.Company{}

	givenDeleteRepoError := errors.New("DELETE REPO ERROR")
	expectedError := util.RepoDeleteErr(errors.New("DELETE REPO ERROR"))

	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.Id)
	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	suite.repo.On("Delete", mock.Anything, givenCompanyIDFilter).Once().Return(givenDeleteRepoError)
	_, err := suite.service.Delete(suite.ctx, givenInput)

	suite.Error(err)
	suite.Equal(err, expectedError)
}
