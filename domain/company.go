package domain

type Company struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}
