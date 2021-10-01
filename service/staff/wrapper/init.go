package wrapper

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type wrapper struct {
	service staff.Service
}

func WrapperStaff(service staff.Service) staff.Service {
	return &wrapper{
		service: service,
	}
}
