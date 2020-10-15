package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
)

type wrapper struct {
	service company.Service
}

func _(service company.Service) company.Service {
	return &wrapper{
		service: service,
	}
}
