package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
)

func (impl *implementation) Update(ctx context.Context, input *pb.UpdateStaffRequest) (item *pb.UpdateStaffResponse, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, nil
		//return util.ValidationUpdateErr(err), nil
	}

	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.Id)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, nil
		//return util.RepoReadErr(err)
	}

	update := staffin.UpdateInputGrpcToStaffInputDomain(input)
	staff.ID = update.ID
	staff.CompanyID = update.CompanyID
	staff.Name = update.Name
	staff.Tel = update.Tel
	staff.UpdatedAt = update.UpdatedAt

	err = impl.repo.Update(ctx, filters, staff)
	if err != nil {
		return nil, nil
		//return util.RepoUpdateErr(err)
	}

	return new(pb.UpdateStaffResponse), nil
}
