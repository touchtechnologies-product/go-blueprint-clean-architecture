package staff

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
)

func (impl *implement) CreateStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := pb.NewStaffGrpcServiceClient(client)

	input := &pb.CreateStaffRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(pb.CreateStaffResponse)})
		return
	}

	output, err := grpcCompanyClient.Create(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(pb.CreateStaffResponse)})
		return
	}

	if output != nil {
		c.JSON(201, gin.H{"success": true, "code": 201, "msg": "Created", "data": output})
		return
	}

	c.JSON(500, gin.H{"success": true, "code": 500, "msg": "Something went wrong!", "data": new(pb.CreateStaffResponse)})
	return
}

func (impl *implement) ListStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := protobuf.NewCompanyGrpcServiceClient(client)

	input := &protobuf.ListCompanyRequest{}
	if err := c.ShouldBindQuery(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(protobuf.ListCompanyResponse)})
		return
	}

	output, err := grpcCompanyClient.List(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(protobuf.ListCompanyResponse)})
		return
	} else {

	}

	if output != nil {
		c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(protobuf.ListCompanyResponse)})
	return
}

func (impl *implement) ReadStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := protobuf.NewCompanyGrpcServiceClient(client)

	input := &protobuf.UpdateCompanyRequest{}
	output, err := grpcCompanyClient.Update(ctx, input)
	if err != nil {
		c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": new(protobuf.UpdateCompanyResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
	return
}

func (impl *implement) UpdateStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := protobuf.NewCompanyGrpcServiceClient(client)

	input := &protobuf.UpdateCompanyRequest{Id: c.Param("id")}

	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": new(protobuf.UpdateCompanyResponse)})
		return
	}

	output, err := grpcCompanyClient.Update(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(protobuf.UpdateCompanyResponse)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Updated", "data": output})
	return
}

func (impl *implement) DeleteStaff(c *gin.Context) {
	ctx := context.Background()
	client := impl.conn
	grpcCompanyClient := protobuf.NewCompanyGrpcServiceClient(client)

	input := &protobuf.DeleteCompanyRequest{Id: c.Param("id")}

	_, err := grpcCompanyClient.Delete(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": new(protobuf.DeleteCompanyRequest)})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "", "data": new(protobuf.DeleteCompanyRequest)})
	return
}
