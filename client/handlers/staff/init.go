package staff

import (
	"google.golang.org/grpc"
)

type implement struct {
	conn *grpc.ClientConn
}

func NewStaffHandler(conn *grpc.ClientConn) (service Handler) {
	return &implement{
		conn,
	}
}
