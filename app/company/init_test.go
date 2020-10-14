package company

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/mocks"
)

type CompanyTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	company *Company
	service *mocks.Service
}

func (suite *CompanyTestSuite) SetupSuite() {
	suite.service = &mocks.Service{}
	suite.company = New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())
	suite.router.POST("/api/v1/company", suite.company.CreateCompany)
}

func TestFormTestSuite(t *testing.T) {
	suite.Run(t, new(CompanyTestSuite))
}
