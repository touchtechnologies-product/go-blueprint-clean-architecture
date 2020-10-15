package test

import (
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	conf    *config.Config
	ctrl    *company.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = company.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.GET("/companies", suite.ctrl.List)
	suite.router.POST("/companies", suite.ctrl.Create)
	suite.router.GET("/companies/:id", suite.ctrl.Read)
	suite.router.PUT("/companies/:id", suite.ctrl.Update)
	suite.router.DELETE("/companies/:id", suite.ctrl.Delete)
}
