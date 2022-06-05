package mock

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type LocationMockRepo struct {
	mock.Mock
	Loc  *model.Location
	Locs []*model.Location
}

func (r *LocationMockRepo) FindOne(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *LocationMockRepo) FindMulti(ids []uint32, locs *[]*model.Location) error {
	args := r.Called(ids, locs)
	if args.Get(0) != nil {
		*locs = args.Get(0).([]*model.Location)
	}

	return args.Error(1)
}

func (r *LocationMockRepo) Create(loc *model.Location) error {
	args := r.Called(loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *LocationMockRepo) Update(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}

func (r *LocationMockRepo) Delete(id int, loc *model.Location) error {
	args := r.Called(id, loc)
	if args.Get(0) != nil {
		*loc = *args.Get(0).(*model.Location)
	}

	return args.Error(1)
}
