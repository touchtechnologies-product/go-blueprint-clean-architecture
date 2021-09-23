package test

import (
	"github.com/stretchr/testify/suite"
	"github.com/uniplaces/carbon"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type PackageTestSuite struct {
	suite.Suite
}

func MakeTestStaff() (staff *domain.Staff) {
	return &domain.Staff{
		ID:        "test",
		CompanyID: "test",
		Name:      "test",
		Tel:       "test",
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}
}
