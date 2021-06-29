package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (suite *PackageTestSuite) TestReadCompanyGPRCSuccess() {
	givenInput := companyin.MakeTestReadGRPCInput()
	givenCompanyIDFilter := suite.makeCompanyIDFilter(givenInput.CompanyId)
	givenCompany := &domain.Company{}
	suite.repo.On("Read", mock.Anything, givenCompanyIDFilter, givenCompany).Once().Return(nil)
	actualView, err := suite.service.Read(suite.ctx, givenInput)
	suite.NoError(err)
	suite.Equal(givenCompany.ID, actualView.Id)
	suite.Equal(givenCompany.Name, actualView.Name)
}

func (suite *PackageTestSuite) TestReadCompanyGPRCError() {
	company := &domain.Company{}
	givenInput := companyin.MakeTestReadGRPCInput()
	filters := suite.makeCompanyIDFilter(givenInput.CompanyId)

	givenReadRepoError := errors.New("READ REPOSITORY")
	expectedError := util.RepoReadErr(errors.New("READ REPOSITORY"))

	suite.repo.On("Read", mock.Anything, filters, company).Once().Return(givenReadRepoError)
	_, err := suite.service.Read(suite.ctx, givenInput)

	suite.Error(err)
	suite.Equal(err, expectedError)
}
