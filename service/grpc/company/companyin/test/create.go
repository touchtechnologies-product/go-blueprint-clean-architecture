package test

import "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

func (suite *PackageTestSuite) TestCreateInputToCompanyDomain() {
	given := companyin.MakeTestCreateInput()

	actual := companyin.CreateInputToCompanyDomain(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
