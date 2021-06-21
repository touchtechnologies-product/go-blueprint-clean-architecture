package company

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

func (impl *implement) CreateCompany(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.CreateCompanyRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.CreateCompanyResponse)})
		return
	}

	output, err := grpcCompanyClient.Create(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.CreateCompanyResponse)})
		return
	}

	if output != nil {
		c.JSON(201, gin.H{"success": true, "code": 201, "msg": "Created", "data": output})
		return
	}

	c.JSON(500, gin.H{"success": true, "code": 500, "msg": "Something went wrong!", "data": new(pb.CreateCompanyResponse)})
	return
}

func (impl *implement) ListCompany(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.ListCompanyRequest{}
	if err := c.ShouldBindQuery(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.ListCompanyResponse)})
		return
	}

	output, err := grpcCompanyClient.List(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.ListCompanyResponse)})
		return
	} else {

	}

	if output != nil {
		c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(pb.ListCompanyResponse)})
	return
}

func (impl *implement) ReadCompany(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.ReadCompanyRequest{CompanyId: c.Param("id")}
	output, err := grpcCompanyClient.Read(ctx, input)
	if err != nil {
		c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(pb.ReadCompanyResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
	return
}

func (impl *implement) UpdateCompany(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.UpdateCompanyRequest{Id: c.Param("id")}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.UpdateCompanyResponse)})
		return
	}

	output, err := grpcCompanyClient.Update(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.UpdateCompanyResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Updated", "data": output})
	return
}

func (impl *implement) DeleteCompany(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.DeleteCompanyRequest{Id: c.Param("id")}

	_, err := grpcCompanyClient.Delete(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.DeleteCompanyRequest)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Deleted", "data": new(pb.DeleteCompanyRequest)})
	return
}
