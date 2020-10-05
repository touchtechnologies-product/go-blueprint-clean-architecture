package company

import (
	"blueprint/service/company"
)

type Company struct {
	service company.Service
}

func New(companyService company.Service) (company *Company) {
	return &Company{service: companyService}
}
