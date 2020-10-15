package implement

import (
	"context"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Update(ctx context.Context, input *staffin.UpdateInput) (err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return util.ValidationUpdateErr(err)
	}

	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.ID)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return util.RepoReadErr(err)
	}

	update := staffin.UpdateInputToStaffDomain(input)
	staff.ID = update.ID
	staff.CompanyID = update.CompanyID
	staff.Name = update.Name
	staff.Tel = update.Tel
	staff.UpdatedAt = update.UpdatedAt

	err = impl.repo.Update(ctx, filters, staff)
	if err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
