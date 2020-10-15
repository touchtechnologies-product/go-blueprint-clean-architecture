package staff_test

import (
	"testing"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
