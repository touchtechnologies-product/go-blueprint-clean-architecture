package view

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"

	"github.com/gin-gonic/gin"
)

const errorStatus = "error"

type errResp struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Errors []*util.Error `json:"errors"`
}

func MakeErrResp(c *gin.Context, err error) {
	errResp := &errResp{
		Status: errorStatus,
		Code:   getHTTPStatusCode(err),
		Errors: getRespErrors(err),
	}
	c.JSON(errResp.Code, errResp)
}

func getHTTPStatusCode(err error) (code int) {
	switch err := util.TypeOfErr(err); {
	case err.IsType(util.ConvertInputToDomainErr):
		return http.StatusBadRequest
	case err.IsType(util.ValidationCreateErr):
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

func getRespErrors(err error) (errs []*util.Error) {
	switch err.(type) {
	case *util.Error:
		return utilToResp(err.(*util.Error))
	default:
		return []*util.Error{util.UnknownErr(err)}
	}
}

func utilToResp(err error) (errs []*util.Error) {
	switch err := util.TypeOfErr(err); {
	case err.IsType(util.ValidationCreateErr):
		return validateToResp(err)
	default:
		return []*util.Error{err}
	}
}

func validateToResp(err *util.Error) (errs []*util.Error) {
	vErrs := err.Cause.(validator.ValidationErrors)
	errs = make([]*util.Error, len(vErrs))
	for i, vErr := range vErrs {
		errs[i] = &util.Error{
			Cause:   errors.New(vErr.Translate(nil)),
			Code:    vErr.Tag(),
			SubCode: vErr.Field(),
		}
	}

	return
}
