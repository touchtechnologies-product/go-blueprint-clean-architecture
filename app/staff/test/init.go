package test

import (
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/staff"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type PackageTestSuite struct {
	suite.Suite
	router  *gin.Engine
	ctx     *gin.Context
	conf    *config.Config
	ctrl    *staff.Controller
	service *mocks.Service
}

func (suite *PackageTestSuite) SetupSuite() {
	suite.conf = config.Get()
	suite.service = &mocks.Service{}
	suite.ctrl = staff.New(suite.service)
	suite.ctx, suite.router = gin.CreateTestContext(httptest.NewRecorder())

	suite.router.GET("/staffs", suite.ctrl.List)
	suite.router.POST("/staffs", suite.ctrl.Create)
	suite.router.GET("/staffs/:id", suite.ctrl.Read)
	suite.router.PUT("/staffs/:id", suite.ctrl.Update)
	suite.router.DELETE("/staffs/:id", suite.ctrl.Delete)
}
