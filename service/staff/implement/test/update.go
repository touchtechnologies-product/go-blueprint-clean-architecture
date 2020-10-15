package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestUpdate() {
	givenInput := staffin.MakeTestUpdateInput()
	givenStaff := domain.MakeTestStaff()
	givenStaff.CreatedAt = 0
	givenStaffIDFilter := suite.makeStaffIDFilter(givenInput.ID)

	suite.repo.On("Read", mock.Anything, givenStaffIDFilter, &domain.Staff{}).Once().Return(nil)
	suite.validator.On("Validate", givenInput).Once().Return(nil)
	suite.repo.On("Update", mock.Anything, givenStaffIDFilter, givenStaff).Once().Return(nil)
	err := suite.service.Update(suite.ctx, givenInput)

	suite.NoError(err)
}
