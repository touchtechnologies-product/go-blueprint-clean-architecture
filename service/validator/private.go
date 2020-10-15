package validator

import (
	"context"
	"fmt"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (v *GoPlayGroundValidator) pageOptionFilterValidation(structLV validator.StructLevel, filter string) {
	fragments := strings.Split(filter, ":")
	if len(fragments) < 2 {
		structLV.ReportError(filter, "filters", "filters", "length", "")
		return
	}

	support := map[string]bool{
		"ne":        true,
		"like":      true,
		"eq":        true,
		"eqInt":     true,
		"isNull":    true,
		"isNotNull": true,
		"id":        true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(filter, "filters", "filters", "support", "")
		return
	}

	switch fragments[1] {
	case "isNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	case "isNotNull":
		if len(fragments) == 2 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	default:
		if len(fragments) == 3 {
			return
		}
		structLV.ReportError(filter, "filters", "filters", "length", "")
	}
}

func (v *GoPlayGroundValidator) pageOptionSortValidation(structLV validator.StructLevel, sort string) {
	fragments := strings.Split(sort, ":")
	if len(fragments) < 2 {
		structLV.ReportError(sort, "sorts", "sorts", "length", "")
		return
	}

	support := map[string]bool{
		"asc":  true,
		"desc": true,
	}
	if _, ok := support[fragments[1]]; !ok {
		structLV.ReportError(sort, "sorts", "sorts", "support", "")
		return
	}
}

func (v *GoPlayGroundValidator) checkCompanyIDUnique(ctx context.Context, structLV validator.StructLevel, ID string) (company *domain.Company) {
	if err := v.companyRepo.Read(ctx, []string{fmt.Sprintf("id:eq:%s", ID)}, &domain.Company{}); err == nil {
		structLV.ReportError(ID, "id", "ID", "unique", "")
	}
	return company
}

func (v *GoPlayGroundValidator) checkCompanyNameUnique(ctx context.Context, structLV validator.StructLevel, companyName string) {
	if err := v.companyRepo.Read(ctx, []string{fmt.Sprintf("name:eq:%s", companyName)}, &domain.Company{}); err == nil {
		structLV.ReportError(companyName, "name", "name", "unique", "")
	}
}

func (v *GoPlayGroundValidator) checkCompanyNameUniqueUpdate(ctx context.Context, structLV validator.StructLevel, input *companyin.UpdateInput) {
	filters := []string{
		fmt.Sprintf("id:ne:%s", input.ID),
		fmt.Sprintf("name:eq:%s", input.Name),
	}
	if err := v.companyRepo.Read(ctx, filters, &domain.Company{}); err == nil {
		structLV.ReportError(input.Name, "name", "name", "unique", "")
	}
}