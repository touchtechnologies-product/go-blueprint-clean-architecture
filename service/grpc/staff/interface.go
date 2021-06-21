package staff

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, input *protobuf.ListStaffRequest) (items *protobuf.ListStaffResponse, err error)
	Create(ctx context.Context, input *protobuf.CreateStaffRequest) (id *protobuf.CreateStaffResponse, err error)
	Read(ctx context.Context, input *protobuf.ReadStaffRequest) (company *protobuf.ReadStaffResponse, err error)
	Update(ctx context.Context, input *protobuf.UpdateStaffRequest) (item *protobuf.UpdateStaffResponse, err error)
	Delete(ctx context.Context, input *protobuf.DeleteStaffRequest) (item *protobuf.DeleteStaffResponse, err error)
}
