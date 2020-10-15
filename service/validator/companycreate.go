package validator

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) CompanyCreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	company := structLV.Current().Interface().(companyin.CreateInput)

	v.checkCompanyIDUnique(ctx, structLV, company.ID)
	v.checkCompanyNameUnique(ctx, structLV, company.Name)
}
