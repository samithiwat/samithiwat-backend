package team

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindAll(pagination *model.TeamPagination) error {
	args := r.Called(pagination)

	if args.Get(0) != nil {
		*pagination = args.Get(0).(model.TeamPagination)
	}

	return args.Error(1)
}

func (r *MockRepo) FindOne(id uint, team *model.Team) error {
	args := r.Called(id, team)

	if args.Get(0) != nil {
		*team = *args.Get(0).(*model.Team)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(ids []uint32, teams *[]*model.Team) error {
	args := r.Called(ids, teams)

	if args.Get(0) != nil {
		*teams = args.Get(0).([]*model.Team)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(team *model.Team) error {
	args := r.Called(team)

	if args.Get(0) != nil {
		*team = *args.Get(0).(*model.Team)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id uint, team *model.Team) error {
	args := r.Called(id, team)

	if args.Get(0) != nil {
		*team = *args.Get(0).(*model.Team)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id uint, team *model.Team) error {
	args := r.Called(id, team)

	if args.Get(0) != nil {
		*team = *args.Get(0).(*model.Team)
	}

	return args.Error(1)
}
