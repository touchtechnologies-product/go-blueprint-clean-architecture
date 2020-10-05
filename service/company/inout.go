package company

import (
	domain "blueprint/domain/company"
	"blueprint/service/util"
)

type CreateInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Paginator struct {
	Items    []*domain.Company
	LastPage int
	*util.Paginator
}

func createInputDomain(input *CreateInput) (staff *domain.Company) {
	return &domain.Company{
		ID:   input.ID,
		Name: input.Name,
	}
}
