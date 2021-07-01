package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/companyin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *pb.CreateCompanyRequest) (output *pb.CreateCompanyResponse, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, util.ValidationCreateErr(err)
	}

	input.Id = impl.uuid.Generate()
	company := companyin.CreateCompGRPCInputToCompanyInputDomain(input)

	_, err = impl.repo.Create(ctx, company)
	if err != nil {
		return nil, util.RepoCreateErr(err)
	}

	return out.OutputCreateCompanyGrpc(company), nil
}
