package out

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

func ListCompanyToListCompanyGrpcView(total int, company []*protobuf.ListCompanyResponse_Output) *protobuf.ListCompanyResponse {
	return &protobuf.ListCompanyResponse{
		Total: int64(total),
		Items: company,
	}
}

func ListItemToListItemGrpcView(company *domain.Company) *protobuf.ListCompanyResponse_Output {
	return &protobuf.ListCompanyResponse_Output{
		Id:   company.ID,
		Name: company.Name,
	}
}

func SingleItemSingleItemGrpcView(company *domain.Company) *protobuf.ReadCompanyResponse {
	return &protobuf.ReadCompanyResponse{
		Id:   company.ID,
		Name: company.Name,
	}
}
