package staff

import (
	"github.com/uniplaces/carbon"

	domain "blueprint/domain/staff"
)

type CreateInput struct {
	Name      string `json:"name" validator:"required"`
	CompanyID string `json:"companyId" validator:"required"`
	Tel       string `json:"tel"`
}

type View struct {
	Name      string `json:"name"`
	CompanyID string `json:"companyId" validator:"required"`
	Tel       string `json:"tel"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}

func (impl *Staff) createInputToStaffDomain(input *CreateInput, timezone string) (staff *domain.Staff, err error) {
	now, err := carbon.NowInLocation(timezone)
	if err != nil {
		return nil, err
	}

	return &domain.Staff{
		ID:        impl.uuid.Generate(),
		CompanyID: input.CompanyID,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: now.Timestamp(),
		UpdatedAt: now.Timestamp(),
	}, nil
}

func staffToView(staff *domain.Staff) (view *View) {
	return &View{
		Name: staff.Name,
	}
}
