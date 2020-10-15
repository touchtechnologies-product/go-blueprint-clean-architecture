package test

import "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

func (suite *PackageTestSuite) TestUpdateInputToStaffDomain() {
	given := staffin.MakeTestUpdateInput()

	actual := staffin.UpdateInputToStaffDomain(given)

	suite.Equal(given.CompanyID, actual.CompanyID)
	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
	suite.Equal(given.Tel, actual.Tel)
}
