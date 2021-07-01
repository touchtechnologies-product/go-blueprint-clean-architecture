package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

type CreateCompGRPCInput struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
}

func CreateCompGRPCInputToCompanyInputDomain(input *pb.CreateCompanyRequest) *domain.Company {
	return &domain.Company{
		ID:   input.Id,
		Name: input.Name,
	}
}

// FUNCTION(S) FOR TEST

func MakeTestCreateCompGRPCInput() (input *pb.CreateCompanyRequest) {
	return &pb.CreateCompanyRequest{
		Name: "test",
	}
}

func MakeTestCreateCompGRPCCToCompanyDomain() *domain.Company {
	return &domain.Company{Name: "test"}
}
