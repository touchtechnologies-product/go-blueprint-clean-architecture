package domain

type Company struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

func MakeTestCompany() (company *Company) {
	return &Company{
		ID:   "test",
		Name: "test",
	}
}