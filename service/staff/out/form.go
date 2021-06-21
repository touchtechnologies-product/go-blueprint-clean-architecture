package out

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	pb "github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/grpc/staff/protobuf"
)

type StaffView struct {
	ID   string `json:"id"`
	Name string `json:"name"`
} // @Name StaffView

func StaffToView(staff *domain.Staff) (view *StaffView) {
	return &StaffView{
		ID:   staff.ID,
		Name: staff.Name,
	}
}

func OutputCreateStaffGrpc(input *domain.Staff) *pb.CreateStaffResponse {
	return &pb.CreateStaffResponse{
		Id:   input.ID,
		Name: input.Name,
	}
}

func ListStaffToListStaffGrpcView(total int, staff []*pb.ListStaffResponse_Output) *pb.ListStaffResponse {
	return &pb.ListStaffResponse{
		Total: int64(total),
		Items: staff,
	}
}

func ListItemToListItemGrpcView(staff *domain.Staff) *pb.ListStaffResponse_Output {
	return &pb.ListStaffResponse_Output{
		Id:   staff.ID,
		Name: staff.Name,
	}
}

func SingleItemSingleItemGrpcView(staff *domain.Staff) *pb.ReadStaffResponse {
	return &pb.ReadStaffResponse{
		Id:   staff.ID,
		Name: staff.Name,
	}
}
