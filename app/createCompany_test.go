package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/stretchr/testify/mock"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/app/inout/company"
	domainCompany "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/company"
	serviceCompany "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

func buildRequestCreateCompany(mode string, input *serviceCompany.CreateCompanyInput) (*http.Request, *httptest.ResponseRecorder) {
	var req *http.Request
	w := httptest.NewRecorder()

	inputBytes, _ := json.Marshal(input)

	switch mode {
	case "success":
		req, _ = http.NewRequest("POST", "/company", bytes.NewBuffer(inputBytes))
		req.Header.Set("Content-Type", "application/json")
	case "notFound":
		req, _ = http.NewRequest("PUT", "/company", bytes.NewBuffer(inputBytes))
		req.Header.Set("Content-Type", "application/json")
	}

	return req, w
}

func (s *AppTestSuite) Test_CreateCompany() {
	expectResponse := &company.CreateCompanyOutput{
		Company: &company.Company{Id: "test_1", Name: "CompanyTest"},
	}

	input := &serviceCompany.CreateCompanyInput{Name: "CompanyTest"}
	req, resp := buildRequestCreateCompany("success", input)

	s.companyService.On("CreateCompany", mock.Anything, &serviceCompany.CreateCompanyInput{Name: input.Name}).Return(&domainCompany.Company{
		Id:   "test_1",
		Name: "CompanyTest",
	}, nil)

	s.router.ServeHTTP(resp, req)

	respData := &company.CreateCompanyOutput{}
	err := json.Unmarshal(resp.Body.Bytes(), respData)

	s.NoError(err)
	s.Equal(http.StatusOK, resp.Code)
	s.Equal(expectResponse, respData)
}

func replaceResponse(bytesBody []byte) string {
	return strings.Replace(string(bytesBody), "\n", "", -1)
}

func (s *AppTestSuite) Test_CreateCompany_InvalidRequest() {
	input := &serviceCompany.CreateCompanyInput{Name: ""}
	req, resp := buildRequestCreateCompany("success", input)

	errorJsonString := `{"errors":[],"message":"invalid request","type":"InvalidRequest"}`

	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusBadRequest, resp.Code)
	s.Equal(errorJsonString, replaceResponse(resp.Body.Bytes()))
}

func (s *AppTestSuite) Test_CreateCompany_MethodNotFound() {
	input := &serviceCompany.CreateCompanyInput{Name: "CompanyTest"}
	req, resp := buildRequestCreateCompany("notFound", input)

	s.router.ServeHTTP(resp, req)

	s.Equal(http.StatusNotFound, resp.Code)
}
