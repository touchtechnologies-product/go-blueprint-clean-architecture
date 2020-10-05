package util

import (
	validator "github.com/go-playground/validator/v10"
)

//go:generate mockery -name=Validator
type Validator interface {
	Validate(item interface{}) (err error)
}

type GoPlayGroundValidator struct {
	validate *validator.Validate
}

func NewValidator() (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{validate: validator.New()}
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
