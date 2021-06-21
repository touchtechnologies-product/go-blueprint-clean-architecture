package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

func (impl *implementation) Delete(ctx context.Context, input *pb.DeleteCompanyRequest) (item *pb.DeleteCompanyResponse, err error) {
	company := &domain.Company{}
	filters := makeCompanyIDFilters(input.Id)

	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, err
	}

	err = impl.repo.Delete(ctx, filters)
	if err != nil {
		return nil, err
	}

	return new(pb.DeleteCompanyResponse), err
}
