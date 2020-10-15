package out

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type CompanyView struct {
	ID    string                      `json:"id"`
	Name  string                      `json:"name"`
} // @Name CompanyView

func CompanyToView(company *domain.Company) (view *CompanyView) {
	return &CompanyView{
		ID:    company.ID,
		Name:  company.Name,
	}
}
