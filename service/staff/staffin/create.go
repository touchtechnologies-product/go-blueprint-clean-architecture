package staffin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
	"github.com/uniplaces/carbon"
	"strconv"
)

type CreateInput struct {
	CompanyID string `json:"company_id" validate:"required"`
	ID        string `json:"id"`
	Name      string `json:"name" validate:"required"`
	Tel       string `json:"tel" validate:"required"`
} // @Name StaffCreateInput

func MakeTestCreateInput() (input *CreateInput) {
	return &CreateInput{
		CompanyID: "test",
		ID:        "test",
		Name:      "test",
		Tel:       "test",
	}
}

func CreateInputToStaffDomain(input *CreateInput) (staff *domain.Staff) {
	return &domain.Staff{
		CompanyID: input.CompanyID,
		ID:        input.ID,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}
}

func CreateInputGrpcToStaffInputDomain(input *pb.CreateStaffRequest) *domain.Staff {
	return &domain.Staff{
		CompanyID: input.CompanyId,
		ID:        input.Id,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}
}

func CreateInputPageOptGrpcToPageOtpDomain(opt *pb.ListStaffRequest) *domain.PageOption {
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
