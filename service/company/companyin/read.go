package companyin

type ReadInput struct {
	CompanyID string `json:"-" validate:"required"`
}

func MakeTestReadInput() (input *ReadInput) {
	return &ReadInput{
		CompanyID: "test",
	}
}
