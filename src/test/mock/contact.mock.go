package mock

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/stretchr/testify/mock"
)

type ContactMockRepo struct {
	mock.Mock
	Cont  *model.Contact
	Conts []*model.Contact
}

func (r *ContactMockRepo) FindOne(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *ContactMockRepo) FindMulti(ids []uint32, conts *[]*model.Contact) error {
	args := r.Called(ids, conts)
	if args.Get(0) != nil {
		*conts = args.Get(0).([]*model.Contact)
	}

	return args.Error(1)
}

func (r *ContactMockRepo) Create(cont *model.Contact) error {
	args := r.Called(cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *ContactMockRepo) Update(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}

func (r *ContactMockRepo) Delete(id int, cont *model.Contact) error {
	args := r.Called(id, cont)
	if args.Get(0) != nil {
		*cont = *args.Get(0).(*model.Contact)
	}

	return args.Error(1)
}
