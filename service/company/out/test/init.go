package test

import (
	"github.com/stretchr/testify/suite"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type PackageTestSuite struct {
	suite.Suite
}

func MakeTestCompany() (company *domain.Company) {
	return &domain.Company{
		ID:   "test",
		Name: "test",
	}
}
