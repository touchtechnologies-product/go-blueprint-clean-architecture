package implement

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/validator"
)

type implementation struct {
	validator validator.Validator
	repo      util.Repository
	uuid      util.UUID
}

func New(validator validator.Validator, repo util.Repository, uuid util.UUID) (service company.Service) {
	return &implementation{validator, repo, uuid}
}
