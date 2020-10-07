package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type View struct {
	Name string `json:"name"`
}

func companyToView(company *domain.Company) (view *View) {
	return &View{
		Name: company.Name,
	}
}
