package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/protobuf"
	"strconv"
)

type CreateInput struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
} // @Name CompanyCreateInput

func MakeTestCreateInput() (input *CreateInput) {
	return &CreateInput{
		ID:   "test",
		Name: "test",
	}
}

func CreateInputToCompanyDomain(input *CreateInput) (company *domain.Company) {
	return &domain.Company{
		ID:   input.ID,
		Name: input.Name,
	}
}

func CreateInputGrpcToCompanyInputDomain(input *pb.CreateCompanyRequest) *domain.Company {
	return &domain.Company{
		ID:   input.Id,
		Name: input.Name,
	}
}

func CreateInputPageOptGrpcToPageOtpDomain(opt *pb.ListCompanyRequest) *domain.PageOption {
	page, err := strconv.Atoi(opt.Page)
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(opt.PerPage)
	if err != nil {
		perPage = 20
	}

	return &domain.PageOption{
		Page:    page,
		PerPage: perPage,
		Filters: opt.Filters,
		Sorts:   opt.Sorts,
	}
}
