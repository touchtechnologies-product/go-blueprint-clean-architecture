package implement

import (
	"context"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/staff/staffin"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *pb.CreateStaffRequest) (output *pb.CreateStaffResponse, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, util.ValidationCreateErr(err)
	}

	input.Id = impl.uuid.Generate()
	staff := staffin.CreateInputGrpcToStaffInputDomain(input)

	_, err = impl.repo.Create(ctx, staff)
	if err != nil {
		return nil, util.RepoCreateErr(err)
	}

	return out.OutputCreateStaffGrpc(staff), nil
}
