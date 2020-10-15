package test

import (
	"errors"
	"net/http"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/view"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

type ValidatorTestStruct struct {
	Title string `validate:"required" json:"title"`
	Body  string `validate:"required"`
}

func (suite *PackageTestSuite) TestMakeErrResp() {
	err := util.ConvertInputToDomainErr(errors.New("test"))
	view.MakeErrResp(suite.ctx, err)
	suite.Equal(http.StatusBadRequest, suite.ctx.Writer.Status())
}
