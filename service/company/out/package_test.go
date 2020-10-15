package out_test

import (
	"testing"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
