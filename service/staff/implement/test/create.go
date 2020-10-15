package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	givenInput := staffin.MakeTestCreateInput()
	givenStaff := domain.MakeTestStaff()

	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Create", mock.Anything, givenStaff).Once().Return(givenStaff.ID, nil)
	actualID, err := suite.service.Create(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenStaff.ID, actualID)
}
