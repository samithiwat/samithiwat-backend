package location_test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/location"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type LocationServiceTest struct {
	suite.Suite
	Loc                   *model.Location
	Locs                  []*model.Location
	CreateLocationReqMock *proto.CreateLocationRequest
	UpdateLocationReqMock *proto.UpdateLocationRequest
}

func TestLocationService(t *testing.T) {
	suite.Run(t, new(LocationServiceTest))
}

func (t *LocationServiceTest) SetupTest() {
	t.Loc = &model.Location{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Address:  faker.Sentence(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc2 := &model.Location{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Address:  faker.Sentence(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc3 := &model.Location{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Address:  faker.Sentence(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	Loc4 := &model.Location{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Address:  faker.Sentence(),
		District: faker.Word(),
		Province: faker.Word(),
		Country:  faker.Word(),
		ZipCode:  faker.Word(),
	}

	t.Locs = append(t.Locs, t.Loc, Loc2, Loc3, Loc4)

	t.CreateLocationReqMock = &proto.CreateLocationRequest{
		Location: &proto.Location{
			Address:  t.Loc.Address,
			District: t.Loc.District,
			Province: t.Loc.Province,
			Country:  t.Loc.Country,
			Zipcode:  t.Loc.ZipCode,
		},
	}

	t.UpdateLocationReqMock = &proto.UpdateLocationRequest{
		Location: &proto.Location{
			Id:       uint32(t.Loc.ID),
			Address:  t.Loc.Address,
			District: t.Loc.District,
			Province: t.Loc.Province,
			Country:  t.Loc.Country,
			Zipcode:  t.Loc.ZipCode,
		},
	}
}

func (t *LocationServiceTest) TestFindOneLocation() {
	var errs []string

	want := &proto.LocationResponse{
		Data:       test.RawToDtoLocation(t.Loc),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &location.MockRepo{}

	r.On("FindOne", 1, &model.Location{}).Return(t.Loc, nil)

	locService := service.NewLocationService(r)
	locRes, err := locService.FindOne(test.Context{}, &proto.FindOneLocationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestFindOneErrNotFoundLocation() {
	errs := []string{"Not found location"}

	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &location.MockRepo{}

	r.On("FindOne", 1, &model.Location{}).Return(nil, errors.New("Not found location"))

	locService := service.NewLocationService(r)
	locRes, err := locService.FindOne(test.Context{}, &proto.FindOneLocationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestFindMultiLocation() {
	var result []*proto.Location
	for _, loc := range t.Locs {
		result = append(result, test.RawToDtoLocation(loc))
	}

	var errs []string
	want := &proto.LocationListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &location.MockRepo{}

	var locs []*model.Location

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &locs).Return(t.Locs, nil)

	locService := service.NewLocationService(r)
	locRes, err := locService.FindMulti(test.Context{}, &proto.FindMultiLocationRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestCreateLocation() {
	var errs []string
	want := &proto.LocationResponse{
		Data:       test.RawToDtoLocation(t.Loc),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	r := &location.MockRepo{}

	locIn := &model.Location{
		Address:  t.Loc.Address,
		District: t.Loc.District,
		Province: t.Loc.Province,
		Country:  t.Loc.Country,
		ZipCode:  t.Loc.ZipCode,
	}

	r.On("Create", locIn).Return(t.Loc, nil)

	locService := service.NewLocationService(r)
	locRes, err := locService.Create(test.Context{}, t.CreateLocationReqMock)

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestUpdateLocation() {
	var errs []string
	want := &proto.LocationResponse{
		Data:       test.RawToDtoLocation(t.Loc),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &location.MockRepo{}

	r.On("Update", 1, t.Loc).Return(t.Loc, nil)

	locService := service.NewLocationService(r)
	locRes, err := locService.Update(test.Context{}, t.UpdateLocationReqMock)

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestUpdateErrNotFoundLocation() {
	errs := []string{"Not found location"}

	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &location.MockRepo{}

	r.On("Update", 1, t.Loc).Return(nil, errors.New("Not found location"))

	locService := service.NewLocationService(r)
	locRes, err := locService.Update(test.Context{}, t.UpdateLocationReqMock)

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestDeleteLocation() {
	var errs []string
	want := &proto.LocationResponse{
		Data:       test.RawToDtoLocation(t.Loc),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &location.MockRepo{}

	r.On("Delete", 1, &model.Location{}).Return(t.Loc, nil)

	locService := service.NewLocationService(r)
	locRes, err := locService.Delete(test.Context{}, &proto.DeleteLocationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}

func (t *LocationServiceTest) TestDeleteErrNotFoundLocation() {
	errs := []string{"Not found location"}

	want := &proto.LocationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &location.MockRepo{}

	r.On("Delete", 1, &model.Location{}).Return(nil, errors.New("Not found location"))

	locService := service.NewLocationService(r)
	locRes, err := locService.Delete(test.Context{}, &proto.DeleteLocationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got any error")
	assert.Equal(t.T(), want, locRes)
}
