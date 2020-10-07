package util

import (
	"context"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type TestStruct struct {
	ID    primitive.ObjectID
	Title string   `validate:"required,min=3"`
	List  []string `validate:"max=1"`
}

type ValidatorTestSuite struct {
	suite.Suite
	ctx       context.Context
	validator *GoPlayGroundValidator
}

func (suite *ValidatorTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.validator = NewValidator()
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

func (suite *ValidatorTestSuite) TestValidateValid() {
	ts := TestStruct{
		Title: "title",
		List:  []string{"a"},
	}
	err := suite.validator.Validate(ts)
	suite.NoError(err)
}

func (suite *ValidatorTestSuite) TestValidateInvalid() {
	ts := TestStruct{
		Title: "",
		List:  []string{"a"},
	}
	err := suite.validator.Validate(ts)
	suite.Error(err)
}
