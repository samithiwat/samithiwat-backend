package location

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindOne(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(ids []uint32, locs *[]*model.Location) error {
	args := r.Called(ids, locs)
	if args.Get(0) != nil {
		*locs = args.Get(0).([]*model.Location)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(loc *model.Location) error {
	args := r.Called(loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}
