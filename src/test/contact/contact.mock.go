package contact

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (r *MockRepo) FindOne(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *MockRepo) FindMulti(ids []uint32, conts *[]*model.Contact) error {
	args := r.Called(ids, conts)
	if args.Get(0) != nil {
		*conts = args.Get(0).([]*model.Contact)
	}

	return args.Error(1)
}

func (r *MockRepo) Create(cont *model.Contact) error {
	args := r.Called(cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *MockRepo) Update(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *MockRepo) Delete(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}
