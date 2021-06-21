package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
)

func (impl *implementation) Delete(ctx context.Context, input *pb.DeleteStaffRequest) (item *pb.DeleteStaffResponse, err error) {
	staff := &domain.Staff{}
	filters := makeStaffIDFilters(input.Id)

	err = impl.repo.Read(ctx, filters, staff)
	if err != nil {
		return nil, err
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return nil, err
	}

	return new(pb.DeleteStaffResponse), err
}
