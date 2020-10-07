package staff

import (
	"blueprint/domain"
	"blueprint/service/util"
	"blueprint/service/util/mocks"
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/uniplaces/carbon"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StaffTestSuite struct {
	suite.Suite
	ctx       context.Context
	repo      *mocks.Repository
	validator *mocks.Validator
	uuid      *mocks.UUID
	svc       Service
}

func (suite *StaffTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.repo = &mocks.Repository{}
	suite.validator = &mocks.Validator{}
	suite.uuid = &mocks.UUID{}
	suite.svc = New(suite.validator, suite.uuid, suite.repo, "Asia/Bangkok")
}

func TestStaffTestSuite(t *testing.T) {
	suite.Run(t, new(StaffTestSuite))
}

func (suite *StaffTestSuite) TestList() {
	listOpt := &util.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: nil,
	}
	suite.repo.On("List", mock.Anything, listOpt, &domain.Staff{}).Once().Return(3, []*domain.Staff{}, nil)
	total, items, err := suite.svc.List(suite.ctx, listOpt)
	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
	suite.Equal(total, 3)
	suite.Equal(items, []*domain.Staff{})
}

func (suite *StaffTestSuite) TestCreate() {
	input := &CreateInput{
		Name:      "Test",
		CompanyID: "Test",
		Tel:       "Test",
	}
	staff := &domain.Staff{
		ID:        "222",
		CompanyID: input.CompanyID,
		Name:      input.Name,
		Tel:       input.Tel,
		CreatedAt: carbon.Now().Timestamp(),
		UpdatedAt: carbon.Now().Timestamp(),
	}

	suite.validator.On("Validate", input).Once().Return(nil)
	suite.uuid.On("Generate").Once().Return("222")
	suite.repo.On("Create", mock.Anything, staff).Once().Return("111", nil)
	ID, err := suite.svc.Create(suite.ctx, input)

	suite.repo.AssertExpectations(suite.T())
	suite.validator.AssertExpectations(suite.T())

	suite.NoError(err)
	suite.Equal(ID, "111")
}

func (suite *StaffTestSuite) TestRead() {
	read := &domain.Staff{}
	ID := "111"

	suite.repo.On("Read", mock.Anything, ID, read).Once().Return(&domain.Staff{})
	expect, err := suite.svc.Read(suite.ctx, ID)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
	suite.Equal(read, expect)
}

func (suite *StaffTestSuite) TestUpdate() {
	input := &CreateInput{
		Name:      "Test",
		CompanyID: "Test",
		Tel:       "Test",
	}
	read := &domain.Staff{}
	ID := "111"

	suite.repo.On("Read", mock.Anything, ID, read).Once().Return(read)
	suite.validator.On("Validate", input).Once().Return(nil)
	suite.repo.On("Update", mock.Anything, ID, read).Once().Return(nil)
	err := suite.svc.Update(suite.ctx, ID, input)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
}

func (suite *StaffTestSuite) TestDelete() {
	read := &domain.Staff{}
	ID := "111"

	suite.repo.On("Read", mock.Anything, ID, read).Once().Return(read)
	suite.repo.On("Delete", mock.Anything, ID).Once().Return(nil)
	err := suite.svc.Delete(suite.ctx, ID)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
}
