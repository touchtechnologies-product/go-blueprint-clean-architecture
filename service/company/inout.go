package company

import (
	domain "blueprint/domain/company"
)

type CreateInput struct {
	Name string `json:"name" validate:"required"`
}

type View struct {
	Name string `json:"name"`
}

func (impl *Company) createInputDomain(input *CreateInput) (staff *domain.Company) {
	return &domain.Company{
		ID:   impl.uuid.Generate(),
		Name: input.Name,
	}
}

func companyToView(company *domain.Company) (view *View) {
	return &View{
		Name: company.Name,
	}
}
