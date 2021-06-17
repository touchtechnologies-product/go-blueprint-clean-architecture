package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/protobuf"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) List(ctx context.Context, optGrpc *pb.ListCompanyRequest) (items *pb.ListCompanyResponse, err error) {

	opt := companyin.CreateInputPageOptGrpcToPageOtpDomain(optGrpc)
	total, records, err := impl.repo.List(ctx, opt, &domain.Company{})
	if err != nil {
		return nil, util.RepoListErr(err)
	}

	output := make([]*pb.ListCompanyResponse_Output, len(records))
	for i, record := range records {
		output[i] = out.ListItemToListItemGrpcView(record.(*domain.Company))
	}

	return out.ListCompanyToListCompanyGrpcView(total, output), nil
}
