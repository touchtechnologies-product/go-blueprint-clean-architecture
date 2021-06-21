package company

import (
	"google.golang.org/grpc"
)

type implement struct {
	conn *grpc.ClientConn
}

func NewCompanyHandler(conn *grpc.ClientConn) (service Handler) {
	return &implement{
		conn,
	}
}
