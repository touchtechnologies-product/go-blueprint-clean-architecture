package test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
)

func (suite *PackageTestSuite) TestCompanyToView() {
	given := domain.MakeTestCompany()

	actual := out.CompanyToView(given)

	suite.Equal(given.ID, actual.ID)
	suite.Equal(given.Name, actual.Name)
}
