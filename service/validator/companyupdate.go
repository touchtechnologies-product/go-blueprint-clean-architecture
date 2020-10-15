package validator

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
)

func (v *GoPlayGroundValidator) CompanyUpdateStructLevelValidation(structLV validator.StructLevel) {
	input := structLV.Current().Interface().(companyin.UpdateInput)
	v.checkCompanyNameUniqueUpdate(context.Background(), structLV, &input)
}
