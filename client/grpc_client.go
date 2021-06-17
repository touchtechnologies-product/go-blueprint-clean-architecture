package main

import "google.golang.org/grpc"

type grpcConn struct {
	conn *grpc.ClientConn
}

func newGrpcClient() (grpcConn *grpc.ClientConn, err error) {
	grpcConn, err = grpc.Dial("localhost:10000", grpc.WithInsecure())
	return
}
