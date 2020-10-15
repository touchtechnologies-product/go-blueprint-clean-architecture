package staffin_test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin/test"
	"testing"


	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
