package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestCreate() {
	input := companyin.MakeTestCreateInput()
	req, resp, err := suite.makeCreateReq(input)
	suite.NoError(err)

	newID := "test"
	suite.service.On("Create", mock.Anything, input).Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
	suite.Equal(newID, resp.Header().Get("Content-Location"))
}

func (suite *PackageTestSuite) makeCreateReq(input *companyin.CreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("POST", "/companies", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}
