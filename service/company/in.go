package company

import (
	"blueprint/domain"
)

type CreateInput struct {
	Name string `json:"name" validate:"required"`
}

func (impl *Company) createInputDomain(input *CreateInput) (staff *domain.Company) {
	return &domain.Company{
		ID:   impl.uuid.Generate(),
		Name: input.Name,
	}
}
