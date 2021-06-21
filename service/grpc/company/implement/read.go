package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/out"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *pb.ReadCompanyRequest) (view *pb.ReadCompanyResponse, err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.CompanyId)
	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	view = out.SingleItemSingleItemGrpcView(company)
	return
}
