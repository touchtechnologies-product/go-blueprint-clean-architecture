package staff

import (
	domain "github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain/staff"
	"github.com/uniplaces/carbon"
)

type CreateInput struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"companyId"`
	Tel       string `json:"tel"`
}

func createInputToStaffDomain(input *CreateInput, timezone string) (staff *domain.Staff, err error) {
	now, err := carbon.NowInLocation(timezone)
	if err != nil {
		return nil, err
	}

	return &domain.Staff{
		Id:        input.Id,
		CompanyId: input.CompanyId,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: now.Timestamp(),
		UpdatedAt: now.Timestamp(),
	}, nil
}