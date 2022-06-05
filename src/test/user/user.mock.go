package user

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindAll(pagination *model.UserPagination) error {
	args := r.Called(pagination)

	if args.Get(0) != nil {
		*pagination = args.Get(0).(model.UserPagination)
	}

	return args.Error(1)
}

func (r *MockRepo) FindOne(id uint, user *model.User) error {
	args := r.Called(id, user)

	if args.Get(0) != nil {
		*user = *args.Get(0).(*model.User)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(ids []uint32, users *[]*model.User) error {
	args := r.Called(ids, users)

	if args.Get(0) != nil {
		*users = args.Get(0).([]*model.User)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(user *model.User) error {
	args := r.Called(user)

	if args.Get(0) != nil {
		*user = *args.Get(0).(*model.User)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id uint, user *model.User) error {
	args := r.Called(id, user)

	if args.Get(0) != nil {
		*user = *args.Get(0).(*model.User)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id uint, user *model.User) error {
	args := r.Called(id, user)

	if args.Get(0) != nil {
		*user = *args.Get(0).(*model.User)
	}

	return args.Error(1)
}
