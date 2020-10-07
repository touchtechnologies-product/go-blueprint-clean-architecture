package staff

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util/mocks"
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
	suite.repo.On("List", mock.Anything, listOpt, domain.Staff{}).Once().Return(3, []interface{}{}, nil)
	total, items, err := suite.svc.List(suite.ctx, listOpt)
	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
	suite.Equal(total, 3)
	suite.Equal(items, []*View{})
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

	suite.repo.On("Read", mock.Anything, []string{"id:eq:111"}, read).Once().Return(nil)
	expect, err := suite.svc.Read(suite.ctx, ID)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
	suite.Equal(expect, &View{})
}

func (suite *StaffTestSuite) TestUpdate() {
	input := &CreateInput{
		Name:      "Test",
		CompanyID: "Test",
		Tel:       "Test",
	}
	read := &domain.Staff{}
	ID := "111"

	suite.repo.On("Read", mock.Anything, []string{"id:eq:111"}, read).Once().Return(nil)
	suite.validator.On("Validate", input).Once().Return(nil)

	update := &domain.Staff{
		Name:      "Test",
		CompanyID: "Test",
		UpdatedAt: carbon.Now().Timestamp(),
	}
	suite.repo.On("Update", mock.Anything, []string{"id:eq:111"}, update).Once().Return(nil)
	err := suite.svc.Update(suite.ctx, ID, input)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
}

func (suite *StaffTestSuite) TestDelete() {
	read := &domain.Staff{}
	ID := "111"

	suite.repo.On("Read", mock.Anything, []string{"id:eq:111"}, read).Once().Return(nil)
	suite.repo.On("Delete", mock.Anything, []string{"id:eq:111"}).Once().Return(nil)
	err := suite.svc.Delete(suite.ctx, ID)

	suite.repo.AssertExpectations(suite.T())
	suite.NoError(err)
}
