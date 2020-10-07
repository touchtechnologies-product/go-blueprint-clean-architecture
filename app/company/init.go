package company

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

type Company struct {
	service company.Service
}

func New(companyService company.Service) (company *Company) {
	return &Company{service: companyService}
}
