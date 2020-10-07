package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
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
