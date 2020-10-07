package view

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	okStatus       = "ok"
	xContentLength = "X-Content-Length"
	location       = "Location"
)

type successResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

func MakeSuccessResp(c *gin.Context, status int, data interface{}) {
	c.JSON(http.StatusOK, successResp{
		Status: okStatus,
		Code:   status,
		Data:   data,
	})
}

func MakePaginatorResp(c *gin.Context, total int, items interface{}) {
	status := http.StatusOK
	if total < 1 {
		status = http.StatusNoContent
	}
	c.Header(xContentLength, strconv.Itoa(total))
	MakeSuccessResp(c, status, items)
}

func MakeCreatedResp(c *gin.Context, ID string) {
	c.Header(location, ID)
	MakeSuccessResp(c, http.StatusCreated, nil)
}
