package company

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

func (suite *CompanyTestSuite) Test_Create() {
	input := &company.CreateInput{
		Name: "Touch",
	}
	req, resp, err := makeCreateReq(input)
	suite.NoError(err)

	newID := "22222"
	suite.service.On("Create", mock.Anything, input).Return(newID, nil)
	suite.router.ServeHTTP(resp, req)

	suite.Equal(http.StatusCreated, resp.Code)
	suite.Equal(newID, resp.Header().Get("Location"))
}
func makeCreateReq(input *company.CreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, nil, err
	}

	req, err = http.NewRequest("POST", "/api/v1/company", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder(), nil
}
