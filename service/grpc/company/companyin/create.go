package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

type GRPCCreateInput struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func MakeTestCreateGRPCInput() (input *pb.CreateCompanyRequest) {
	return &pb.CreateCompanyRequest{
		Name: "test",
	}
}

func MakeTestCreateGRPCCompanyToCompanyDomain() *domain.Company {
	return &domain.Company{Name: "test"}
}

func CreateInputGRPCToCompanyInputDomain(input *pb.CreateCompanyRequest) *domain.Company {
	return &domain.Company{
		ID:   input.Id,
		Name: input.Name,
	}
}
