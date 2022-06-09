package organization

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindAll(pagination *model.OrganizationPagination) error {
	args := r.Called(pagination)

	if args.Get(0) != nil {
		*pagination = args.Get(0).(model.OrganizationPagination)
	}

	return args.Error(1)
}

func (r *MockRepo) FindOne(id uint, org *model.Organization) error {
	args := r.Called(id, org)

	if args.Get(0) != nil {
		*org = *args.Get(0).(*model.Organization)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(ids []uint32, orgs *[]*model.Organization) error {
	args := r.Called(ids, orgs)

	if args.Get(0) != nil {
		*orgs = args.Get(0).([]*model.Organization)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(org *model.Organization) error {
	args := r.Called(org)

	if args.Get(0) != nil {
		*org = *args.Get(0).(*model.Organization)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id uint, org *model.Organization) error {
	args := r.Called(id, org)

	if args.Get(0) != nil {
		*org = *args.Get(0).(*model.Organization)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id uint, org *model.Organization) error {
	args := r.Called(id, org)

	if args.Get(0) != nil {
		*org = *args.Get(0).(*model.Organization)
	}

	return args.Error(1)
}
