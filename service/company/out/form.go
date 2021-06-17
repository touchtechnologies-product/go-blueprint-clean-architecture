package out

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/protobuf"
)

type CompanyView struct {
	ID   string `json:"id"`
	Name string `json:"name"`
} // @Name CompanyView

func CompanyToView(company *domain.Company) (view *CompanyView) {
	return &CompanyView{
		ID:   company.ID,
		Name: company.Name,
	}
}

func OutputCreateCompanyGrpc(input *domain.Company) *pb.CreateCompanyResponse {
	return &pb.CreateCompanyResponse{
		Id:   input.ID,
		Name: input.Name,
	}
}

func ListCompanyToListCompanyGrpcView(total int, company []*pb.ListCompanyResponse_Output) *pb.ListCompanyResponse {
	return &pb.ListCompanyResponse{
		Total: int64(total),
		Items: company,
	}
}

func ListItemToListItemGrpcView(company *domain.Company) *pb.ListCompanyResponse_Output {
	return &pb.ListCompanyResponse_Output{
		Id:   company.ID,
		Name: company.Name,
	}
}

func SingleItemSingleItemGrpcView(company *domain.Company) *pb.ReadCompanyResponse {
	return &pb.ReadCompanyResponse{
		Id:   company.ID,
		Name: company.Name,
	}
}
