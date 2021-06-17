package main

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/protobuf"
	"google.golang.org/grpc"
)

type handler interface {
	createCompany(c *gin.Context)
	listCompany(c *gin.Context)
	readCompany(c *gin.Context)
}

func newCompanyHandler(conn *grpc.ClientConn) (service handler) {
	return &grpcConn{
		conn,
	}
}

func (gc *grpcConn) createCompany(c *gin.Context) {
	ctx := context.Background()
	client := gc.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.CreateCompanyRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": ""})
		return
	}

	output, err := grpcCompanyClient.Create(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": ""})
		return
	}

	if output != nil {
		c.JSON(201, gin.H{"success": true, "code": 201, "msg": "Created", "data": output})
		return
	}

	c.JSON(500, gin.H{"success": true, "code": 500, "msg": "Something went wrong!", "data": ""})
	return
}

func (gc *grpcConn) listCompany(c *gin.Context) {
	ctx := context.Background()
	client := gc.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.ListCompanyRequest{}
	if err := c.ShouldBindQuery(input); err != nil {
		c.JSON(400, gin.H{"success": false, "code": 400, "msg": err.Error(), "data": ""})
		return
	}

	output, err := grpcCompanyClient.List(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": ""})
		return
	}

	if output != nil {
		c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": ""})
	return
}

func (gc *grpcConn) readCompany(c *gin.Context) {
	ctx := context.Background()
	client := gc.conn
	grpcCompanyClient := pb.NewCompanyGrpcServiceClient(client)

	input := &pb.ReadCompanyRequest{CompanyId: c.Param("id")}
	output, err := grpcCompanyClient.Read(ctx, input)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "code": 500, "msg": err.Error(), "data": ""})
		return
	}

	if output != nil {
		c.JSON(200, gin.H{"success": true, "code": 200, "msg": "Successfully", "data": output})
		return
	}

	c.JSON(200, gin.H{"success": true, "code": 404, "msg": "Data not found", "data": ""})
	return
}
