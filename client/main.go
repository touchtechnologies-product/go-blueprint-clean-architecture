package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterRoute(r)
	err := r.Run(":4000")
	if err != nil {
		panic(err)
	}
}

func RegisterRoute(r *gin.Engine) {
	conn, err := newGrpcClient()
	if err != nil {
		panic(err)
	}

	handlers := newCompanyHandler(conn)
	apiRoutes := r.Group("/api/v1")
	{
		apiRoutes.GET("/companies", handlers.listCompany)
		apiRoutes.GET("/companies/:id", handlers.readCompany)
		apiRoutes.POST("/companies", handlers.createCompany)
	}
}
