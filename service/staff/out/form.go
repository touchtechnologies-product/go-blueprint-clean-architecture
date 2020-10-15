package out

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

type StaffView struct {
	ID    string                      `json:"id"`
	Name  string                      `json:"name"`
} // @Name StaffView

func StaffToView(staff *domain.Staff) (view *StaffView) {
	return &StaffView{
		ID:    staff.ID,
		Name:  staff.Name,
	}
}
