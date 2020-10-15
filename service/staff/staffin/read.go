package staffin

type ReadInput struct {
	StaffID string `json:"-" validate:"required"`
}

func MakeTestReadInput() (input *ReadInput) {
	return &ReadInput{
		StaffID: "test",
	}
}
