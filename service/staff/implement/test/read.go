package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRead() {
	givenInput := staffin.MakeTestReadInput()
	givenStaffIDFilter := suite.makeStaffIDFilter(givenInput.StaffID)
	givenStaff := &domain.Staff{}

	suite.repo.On("Read", mock.Anything, givenStaffIDFilter, givenStaff).Once().Return(nil)
	actualView, err := suite.service.Read(suite.ctx, givenInput)

	suite.NoError(err)
	suite.Equal(givenStaff.ID, actualView.ID)
	suite.Equal(givenStaff.Name, actualView.Name)
}
