package companyin

import pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"

type DeleteInput struct {
	ID string `json:"-" validate:"required"`
}

func MakeTestDeleteCompGRPCInput() (input *pb.DeleteCompanyRequest) {
	return &pb.DeleteCompanyRequest{
		Id: "test",
	}
}
