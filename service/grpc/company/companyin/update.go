package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

type GRPCUpdateInput struct {
	Name string `json:"name" validate:"required"`
}

func MakeTestUpdateGRPCInput() (input *pb.UpdateCompanyRequest) {
	return &pb.UpdateCompanyRequest{
		Id:   "test",
		Name: "test",
	}
}

func MakeTestUpdateGRPCCompanyToCompanyDomain() *domain.Company {
	return &domain.Company{ID: "test", Name: "test"}
}

func UpdateInputGRPCToCompanyInputDomain(input *pb.UpdateCompanyRequest) (company *domain.Company) {
	return &domain.Company{
		ID:   input.Id,
		Name: input.Name,
	}
}
