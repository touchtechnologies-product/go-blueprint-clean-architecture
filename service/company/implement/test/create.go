package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := companyin.MakeTestCreateInput()
	givenCompany := suite.makeDataTestCompany

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenCompany).Once().Return(givenCompany.ID, nil)
	actualID, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenCompany.ID, actualID)
}
