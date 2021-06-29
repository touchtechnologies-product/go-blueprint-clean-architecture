package companyin

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type CreateInput struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required"`
} // @Name CompanyCreateInput

func MakeTestCreateInput() (input *CreateInput) {
	return &CreateInput{
		Name: "test",
	}
}

func CreateInputToCompanyDomain(input *CreateInput) (company *domain.Company) {
	return &domain.Company{
		ID:   input.ID,
		Name: input.Name,
	}
}
