package permission

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindAll(pagination *model.PermissionPagination) error {
	args := r.Called(pagination)

	if args.Get(0) != nil {
		*pagination = args.Get(0).(model.PermissionPagination)
	}

	return args.Error(1)
}

func (r *MockRepo) FindOne(id int, perm *model.Permission) error {
	args := r.Called(id, perm)

	if args.Get(0) != nil {
		*perm = *args.Get(0).(*model.Permission)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(perm *model.Permission) error {
	args := r.Called(perm)

	if args.Get(0) != nil {
		*perm = *args.Get(0).(*model.Permission)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id int, perm *model.Permission) error {
	args := r.Called(id, perm)

	if args.Get(0) != nil {
		*perm = *args.Get(0).(*model.Permission)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id int, perm *model.Permission) error {
	args := r.Called(id, perm)

	if args.Get(0) != nil {
		*perm = *args.Get(0).(*model.Permission)
	}

	return args.Error(1)
}
