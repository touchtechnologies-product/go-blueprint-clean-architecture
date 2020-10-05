package staff

import (
	"blueprint/service/staff"
)

type Staff struct {
	service staff.Service
}

func New(staffService staff.Service) (staff *Staff) {
	return &Staff{service: staffService}
}
