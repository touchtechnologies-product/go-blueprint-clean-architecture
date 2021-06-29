package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := companyin.MakeTestCreateInput()
	givenCompany := domain.MakeTestCompany()
	suite.uuid.On("Generate").Once().Return("")
	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenCompany).Once().Return(givenCompany.ID, nil)
	actualID, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenCompany.ID, actualID)
}
