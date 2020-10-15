package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/stretchr/testify/mock"
)

func (suite *PackageTestSuite) TestDelete() {
	input := staffin.MakeTestDeleteInput()
	req, resp, err := suite.makeDeleteReq(input)
	suite.NoError(err)

	suite.service.On("Delete", mock.Anything, input).Return(nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusOK, resp.Code)
}

func (suite *PackageTestSuite) makeDeleteReq(input *staffin.DeleteInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("DELETE", fmt.Sprintf("/staffs/%s", input.ID), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}
