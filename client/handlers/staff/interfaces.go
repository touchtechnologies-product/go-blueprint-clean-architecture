package staff

import "github.com/gin-gonic/gin"

type Handler interface {
	CreateStaff(c *gin.Context)
	ListStaff(c *gin.Context)
	ReadStaff(c *gin.Context)
	UpdateStaff(c *gin.Context)
	DeleteStaff(c *gin.Context)
}
