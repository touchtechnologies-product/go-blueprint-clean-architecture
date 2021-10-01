package wrapper

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

type wrapper struct {
	service company.Service
}

func WrapCompany(service company.Service) company.Service {
	return &wrapper{
		service: service,
	}
}
