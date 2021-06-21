package main

import (
	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/client/grpc_client"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/client/handlers/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/client/handlers/staff"
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
	conn, err := grpc_client.NewGrpcClient()
	if err != nil {
		panic(err)
	}

	companyHandlers := company.NewCompanyHandler(conn)
	staffHandlers := staff.NewStaffHandler(conn)

	apiRoutes := r.Group("/api/v1")
	{
		apiRoutes.GET("/companies", companyHandlers.ListCompany)
		apiRoutes.GET("/companies/:id", companyHandlers.ReadCompany)
		apiRoutes.POST("/companies", companyHandlers.CreateCompany)
		apiRoutes.PUT("/companies/:id", companyHandlers.UpdateCompany)
		apiRoutes.DELETE("/companies/:id", companyHandlers.DeleteCompany)

		apiRoutes.GET("/staff", staffHandlers.ListStaff)
		apiRoutes.GET("/staff/:id", staffHandlers.ReadStaff)
		apiRoutes.POST("/staff", staffHandlers.CreateStaff)
		apiRoutes.PUT("/staff/:id", staffHandlers.UpdateStaff)
		apiRoutes.DELETE("/staff/:id", staffHandlers.DeleteStaff)
	}
}
