package grpc_client

import (
	"google.golang.org/grpc"
)

func NewGrpcClient() (grpcConn *grpc.ClientConn, err error) {
	grpcConn, err = grpc.Dial("localhost:10000", grpc.WithInsecure())
	return
}
