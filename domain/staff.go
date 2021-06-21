package domain

import "github.com/uniplaces/carbon"

type Staff struct {
	ID        string `bson:"_id"`
	CompanyID string `bson:"companyID"`
	Name      string `bson:"name"`
	Tel       string `bson:"tel"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}

func MakeTestStaff() (staff *Staff) {
	return &Staff{
		ID:        "test",
		CompanyID: "test",
		Name:      "test",
		Tel:       "test",
		CreatedAt: carbon.Now().Unix(),
		UpdatedAt: carbon.Now().Unix(),
	}
}
