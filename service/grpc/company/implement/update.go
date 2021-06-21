package implement

import (
	"context"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/company/companyin"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

func (impl *implementation) Update(ctx context.Context, input *pb.UpdateCompanyRequest) (item *pb.UpdateCompanyResponse, err error) {
	err = impl.validator.Validate(input)
	if err != nil {
		return nil, err
	}

	filters := makeCompanyIDFilters(input.Id)

	company := &domain.Company{}
	err = impl.repo.Read(ctx, filters, company)
	if err != nil {
		return nil, err
	}

	update := companyin.UpdateInputGrpcToCompanyInputDomain(input)
	company.Name = update.Name

	err = impl.repo.Update(ctx, filters, company)
	if err != nil {
		return nil, err
	}

	return new(pb.UpdateCompanyResponse), nil
}
