package staff

import (
	"github.com/uniplaces/carbon"

	"blueprint/domain"
)

type CreateInput struct {
	Name      string `json:"name" validator:"required"`
	CompanyID string `json:"companyId" validator:"required"`
	Tel       string `json:"tel"`
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
