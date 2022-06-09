package role

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindAll(pagination *model.RolePagination) error {
	args := r.Called(pagination)

	if args.Get(0) != nil {
		*pagination = args.Get(0).(model.RolePagination)
	}

	return args.Error(1)
}

func (r *MockRepo) FindOne(id int, role *model.Role) error {
	args := r.Called(id, role)

	if args.Get(0) != nil {
		*role = *args.Get(0).(*model.Role)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(id []uint32, roles *[]*model.Role) error {
	args := r.Called(id, roles)

	if args.Get(0) != nil {
		*roles = args.Get(0).([]*model.Role)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(role *model.Role) error {
	args := r.Called(role)

	if args.Get(0) != nil {
		*role = *args.Get(0).(*model.Role)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id int, role *model.Role) error {
	args := r.Called(id, role)

	if args.Get(0) != nil {
		*role = *args.Get(0).(*model.Role)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id int, role *model.Role) error {
	args := r.Called(id, role)

	if args.Get(0) != nil {
		*role = *args.Get(0).(*model.Role)
	}

	return args.Error(1)
}
