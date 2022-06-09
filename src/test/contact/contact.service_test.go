package contact_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/contact"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type ContactServiceTest struct {
	suite.Suite
	Cont                 *model.Contact
	Conts                []*model.Contact
	CreateContactReqMock *proto.CreateContactRequest
	UpdateContactReqMock *proto.UpdateContactRequest
}

func TestContactService(t *testing.T) {
	suite.Run(t, new(ContactServiceTest))
}

func (t *ContactServiceTest) SetupTest() {
	t.Cont = &model.Contact{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Twitter:   faker.URL(),
		Linkedin:  faker.URL(),
	}

	Cont2 := &model.Contact{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Twitter:   faker.URL(),
		Linkedin:  faker.URL(),
	}

	Cont3 := &model.Contact{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Twitter:   faker.URL(),
		Linkedin:  faker.URL(),
	}

	Cont4 := &model.Contact{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Facebook:  faker.URL(),
		Instagram: faker.URL(),
		Twitter:   faker.URL(),
		Linkedin:  faker.URL(),
	}

	t.Conts = append(t.Conts, t.Cont, Cont2, Cont3, Cont4)

	t.CreateContactReqMock = &proto.CreateContactRequest{
		Contact: &proto.Contact{
			Facebook:  t.Cont.Facebook,
			Instagram: t.Cont.Instagram,
			Twitter:   t.Cont.Twitter,
			Linkedin:  t.Cont.Linkedin,
		},
	}

	t.UpdateContactReqMock = &proto.UpdateContactRequest{
		Contact: &proto.Contact{
			Id:        uint32(t.Cont.ID),
			Facebook:  t.Cont.Facebook,
			Instagram: t.Cont.Instagram,
			Twitter:   t.Cont.Twitter,
			Linkedin:  t.Cont.Linkedin,
		},
	}
}

func (t *ContactServiceTest) TestFindOneContact() {
	var errs []string

	want := &proto.ContactResponse{
		Data:       test.RawToDtoContact(t.Cont),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &contact.MockRepo{}

	r.On("FindOne", 1, &model.Contact{}).Return(t.Cont, nil)

	contService := service.NewContactService(r)
	contRes, err := contService.FindOne(test.Context{}, &proto.FindOneContactRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestFindOneErrNotFoundContact() {
	errs := []string{"Not found contact"}

	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &contact.MockRepo{}

	r.On("FindOne", 1, &model.Contact{}).Return(nil, errors.New("Not found contact"))

	contService := service.NewContactService(r)
	contRes, err := contService.FindOne(test.Context{}, &proto.FindOneContactRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestFindMultiContact() {
	var result []*proto.Contact
	for _, cont := range t.Conts {
		result = append(result, test.RawToDtoContact(cont))
	}

	var errs []string

	want := &proto.ContactListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &contact.MockRepo{}

	var conts []*model.Contact

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &conts).Return(t.Conts, nil)

	contService := service.NewContactService(r)
	contRes, err := contService.FindMulti(test.Context{}, &proto.FindMultiContactRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestCreateContact() {
	var errs []string

	want := &proto.ContactResponse{
		Data:       test.RawToDtoContact(t.Cont),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	r := &contact.MockRepo{}

	contIn := &model.Contact{
		Facebook:  t.Cont.Facebook,
		Instagram: t.Cont.Instagram,
		Twitter:   t.Cont.Twitter,
		Linkedin:  t.Cont.Linkedin,
	}

	r.On("Create", contIn).Return(t.Cont, nil)

	contService := service.NewContactService(r)
	contRes, err := contService.Create(test.Context{}, t.CreateContactReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestUpdateContact() {
	var errs []string

	want := &proto.ContactResponse{
		Data:       test.RawToDtoContact(t.Cont),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &contact.MockRepo{}

	r.On("Update", 1, t.Cont).Return(t.Cont, nil)

	contService := service.NewContactService(r)
	contRes, err := contService.Update(test.Context{}, t.UpdateContactReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestUpdateErrNotFoundContact() {
	errs := []string{"Not found contact"}

	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &contact.MockRepo{}

	r.On("Update", 1, t.Cont).Return(nil, errors.New("Not found contact"))

	contService := service.NewContactService(r)
	contRes, err := contService.Update(test.Context{}, t.UpdateContactReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestDeleteContact() {
	var errs []string

	want := &proto.ContactResponse{
		Data:       test.RawToDtoContact(t.Cont),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &contact.MockRepo{}

	r.On("Delete", 1, &model.Contact{}).Return(t.Cont, nil)

	contService := service.NewContactService(r)
	contRes, err := contService.Delete(test.Context{}, &proto.DeleteContactRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}

func (t *ContactServiceTest) TestDeleteErrNotFoundContact() {
	errs := []string{"Not found contact"}

	want := &proto.ContactResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &contact.MockRepo{}

	r.On("Delete", 1, &model.Contact{}).Return(nil, errors.New("Not found contact"))

	contService := service.NewContactService(r)
	contRes, err := contService.Delete(test.Context{}, &proto.DeleteContactRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, contRes)
}
