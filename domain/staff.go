package domain

type Staff struct {
	ID        string `bson:"id"`
	CompanyID string `bson:"companyID"`
	Name      string `bson:"name"`
	Tel       string `bson:"tel"`
	CreatedAt int64  `bson:"createdAt"`
	UpdatedAt int64  `bson:"updatedAt"`
}
