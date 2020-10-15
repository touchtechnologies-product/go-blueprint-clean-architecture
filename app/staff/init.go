package staff

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff"
)

type Controller struct {
	service       staff.Service
}

func New(service staff.Service) (ctrl *Controller) {
	return &Controller{service}
}
