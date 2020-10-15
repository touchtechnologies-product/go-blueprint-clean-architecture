package test

import (
	"github.com/stretchr/testify/mock"

	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/domain"
)

func (suite *PackageTestSuite) TestList() {
	givenPageOption := &domain.PageOption{
		Page:    1,
		PerPage: 10,
		Filters: nil,
	}
	givenTypeGuide := &domain.Staff{}

	givenTotal := 20
	givenItems := make([]interface{}, givenPageOption.PerPage)
	for i := 0; i < givenPageOption.PerPage; i++ {
		givenItems[i] = &domain.Staff{}
	}

	suite.repo.On("List", mock.Anything, givenPageOption, givenTypeGuide).Once().Return(givenTotal, givenItems, nil)
	total, items, err := suite.service.List(suite.ctx, givenPageOption)

	suite.NoError(err)
	suite.Equal(givenTotal, total)
	suite.Equal(len(givenItems), len(items))
}
