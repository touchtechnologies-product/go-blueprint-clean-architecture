package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	"strconv"
)

type CreateInput struct {
	ID   string `json:"id"`
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

func CreateInputGrpcToCompanyInputDomain(input *protobuf.CreateCompanyRequest) *domain.Company {
	return &domain.Company{
		ID:   input.Id,
		Name: input.Name,
	}
}

func CreateInputPageOptGrpcToPageOtpDomain(opt *protobuf.ListCompanyRequest) *domain.PageOption {
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
