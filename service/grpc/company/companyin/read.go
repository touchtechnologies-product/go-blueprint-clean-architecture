package companyin

import (
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/company/protobuf"
)

func MakeTestReadGRPCInput() (input *pb.ReadCompanyRequest) {
	return &pb.ReadCompanyRequest{
		CompanyId: "test",
	}
}
