package staffin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/uniplaces/carbon"
)

type UpdateInput struct {
	CompanyID string `json:"companyID" validate:"required"`
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Tel       string `json:"tel" validate:"required"`
}// @Name StaffUpdateInput

func MakeTestUpdateInput() (input *UpdateInput) {
	return &UpdateInput{
		CompanyID: "test",
		ID:        "test",
		Name:      "test",
		Tel:       "test",
	}
}

func UpdateInputToStaffDomain(input *UpdateInput) (staff *domain.Staff) {
	return &domain.Staff{
		CompanyID: input.CompanyID,
		ID:        input.ID,
		Name:      input.Name,
		Tel:       input.Tel,
		UpdatedAt: carbon.Now().Unix(),
	}
}
