package staff

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, input *protobuf.ListCompanyRequest) (items *protobuf.ListCompanyResponse, err error)
	Create(ctx context.Context, input *protobuf.CreateCompanyRequest) (id *protobuf.CreateCompanyResponse, err error)
	Read(ctx context.Context, input *protobuf.ReadCompanyRequest) (company *protobuf.ReadCompanyResponse, err error)
	Update(ctx context.Context, input *protobuf.UpdateCompanyRequest) (item *protobuf.UpdateCompanyResponse, err error)
	Delete(ctx context.Context, input *protobuf.DeleteCompanyRequest) (item *protobuf.DeleteCompanyResponse, err error)
}
