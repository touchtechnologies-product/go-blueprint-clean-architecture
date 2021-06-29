package implement_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/implement/test"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
