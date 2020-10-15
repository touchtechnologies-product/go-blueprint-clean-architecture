package companyin_test

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin/test"
	"testing"


	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
