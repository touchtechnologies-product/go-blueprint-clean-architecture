// +build integration

package mongodb_test

import (
	"testing"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/repository/mongodb/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
