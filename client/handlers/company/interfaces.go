package company

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateCompany(c *gin.Context)
	ListCompany(c *gin.Context)
	ReadCompany(c *gin.Context)
	UpdateCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
}
