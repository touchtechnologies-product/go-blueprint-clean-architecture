package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *staffin.CreateInput) (ID string, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return "", util.ValidationCreateErr(err)
	}

	staff := staffin.CreateInputToStaffDomain(input)

	_, err = impl.repo.Create(ctx, staff)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return staff.ID, nil
}
