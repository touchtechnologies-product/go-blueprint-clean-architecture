package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type UpdateInput struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}// @Name CompanyUpdateInput

func MakeTestUpdateInput() (input *UpdateInput) {
	return &UpdateInput{
		ID:   "test",
		Name: "test",
	}
}

func UpdateInputToCompanyDomain(input *UpdateInput) (company *domain.Company) {
	return &domain.Company{
		ID:   input.ID,
		Name: input.Name,
	}
}
