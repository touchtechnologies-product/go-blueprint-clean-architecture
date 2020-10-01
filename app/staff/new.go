package staff

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type Staff struct {
	service staff.Service
}

func New(staffService staff.Service) (staff *Staff) {
	return &Staff{service: staffService}
}
