package staffin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/uniplaces/carbon"
)

type CreateInput struct {
	CompanyID string `json:"companyID" validate:"required"`
	ID        string `json:"id" validate:"required"`
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
