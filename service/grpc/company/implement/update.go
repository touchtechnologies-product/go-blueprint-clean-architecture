package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/companyin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

func (impl *implementation) Update(ctx context.Context, input *pb.UpdateCompanyRequest) (item *pb.UpdateCompanyResponse, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, util.ValidationUpdateErr(err)
	}

	filters := makeCompanyIDFilters(input.Id)

	company := &domain.Company{}
	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, util.RepoReadErr(err)
	}

	update := companyin.UpdateInputGRPCToCompanyInputDomain(input)

	err = impl.repo.Update(ctx, filters, update)
	if err != nil {
		return nil, util.RepoUpdateErr(err)
	}

	return new(pb.UpdateCompanyResponse), nil
}
