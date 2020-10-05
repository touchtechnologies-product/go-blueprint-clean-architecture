package staff

import (
	domain "blueprint/domain/staff"
	"blueprint/service/util"

	"github.com/uniplaces/carbon"
)

type CreateInput struct {
	ID        string `json:"id" validator:"required"`
	Name      string `json:"name" validator:"required"`
	CompanyID string `json:"companyId" validator:"required"`
	Tel       string `json:"tel"`
}

type Paginator struct {
	Items    []*domain.Staff
	PageLeft int
	*util.Paginator
}

func createInputToStaffDomain(input *CreateInput, timezone string) (staff *domain.Staff, err error) {
	now, err := carbon.NowInLocation(timezone)
	if err != nil {
		return nil, err
	}

	return &domain.Staff{
		ID:        input.ID,
		CompanyID: input.CompanyID,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: now.Timestamp(),
		UpdatedAt: now.Timestamp(),
	}, nil
}
