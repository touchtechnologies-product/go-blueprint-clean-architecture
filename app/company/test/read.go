package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestRead() {
	input := companyin.MakeTestReadInput()
	req, resp, err := makeReadReq(input)
	suite.NoError(err)

	suite.service.On("Read", mock.Anything, input).Return(&out.CompanyView{}, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func makeReadReq(input *companyin.ReadInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, err = http.NewRequest("GET", fmt.Sprintf("/companies/%s", input.CompanyID), nil)
	if err != nil {
		return nil, nil, err
	}
	return req, httptest.NewRecorder(), nil
}
