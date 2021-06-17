package company

import (
	"context"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/protobuf"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, input *pb.ListCompanyRequest) (items *pb.ListCompanyResponse, err error)
	Create(ctx context.Context, input *pb.CreateCompanyRequest) (id *pb.CreateCompanyResponse, err error)
	Read(ctx context.Context, input *pb.ReadCompanyRequest) (company *pb.ReadCompanyResponse, err error)
	Update(ctx context.Context, input *pb.UpdateCompanyRequest) (item *pb.UpdateCompanyResponse, err error)
	Delete(ctx context.Context, input *pb.DeleteCompanyRequest) (item *pb.DeleteCompanyResponse, err error)
}
