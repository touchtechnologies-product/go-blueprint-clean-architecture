package domain

type Company struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

func MakeTestCompany() (company *Company) {
	return &Company{
		Name: "test",
	}
}
