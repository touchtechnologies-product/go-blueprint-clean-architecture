package view

import (
	"blueprint/service/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const errorStatus = "error"

type ErrResp struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Errors []*util.Error `json:"errors"`
}

func MakeErrResp(c *gin.Context, err error) {
	errResp := &ErrResp{
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
		return http.StatusUnprocessableEntity
	default:
		return []*util.Error{err}
	}
}
