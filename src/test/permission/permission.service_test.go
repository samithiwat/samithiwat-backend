package permission_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/permission"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type PermissionServiceTest struct {
	suite.Suite
	Perm                    *model.Permission
	Perms                   []*model.Permission
	CreatePermissionReqMock *proto.CreatePermissionRequest
	UpdatePermissionReqMock *proto.UpdatePermissionRequest
}

func TestPermissionService(t *testing.T) {
	suite.Run(t, new(PermissionServiceTest))
}

func (t *PermissionServiceTest) SetupTest() {
	t.Perm = &model.Permission{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: faker.Word(),
		Code: faker.Word(),
	}

	Perm2 := &model.Permission{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: faker.Word(),
		Code: faker.Word(),
	}

	Perm3 := &model.Permission{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: faker.Word(),
		Code: faker.Word(),
	}

	Perm4 := &model.Permission{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name: faker.Word(),
		Code: faker.Word(),
	}

	t.Perms = append(t.Perms, t.Perm, Perm2, Perm3, Perm4)

	t.CreatePermissionReqMock = &proto.CreatePermissionRequest{
		Permission: &proto.Permission{
			Name: t.Perm.Name,
			Code: t.Perm.Code,
		},
	}

	t.UpdatePermissionReqMock = &proto.UpdatePermissionRequest{
		Permission: &proto.Permission{
			Id:   uint32(t.Perm.ID),
			Name: t.Perm.Name,
			Code: t.Perm.Code,
		},
	}
}

func (t *PermissionServiceTest) TestFindAllPermission() {
	var result []*proto.Permission
	for _, perm := range t.Perms {
		result = append(result, test.RawToDtoPermission(perm))
	}

	var errs []string

	want := &proto.PermissionListResponse{
		Data: &proto.PermissionPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    int64(len(t.Perms)),
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var perms []*model.Permission

	paginationIn := &model.PermissionPagination{
		Items: &perms,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			CurrentPage:  1,
		},
	}

	paginationOut := model.PermissionPagination{
		Items: &t.Perms,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			ItemCount:    int64(len(t.Perms)),
			TotalItem:    4,
			CurrentPage:  1,
			TotalPage:    1,
		},
	}

	r := &permission.MockRepo{}

	r.On("FindAll", paginationIn).Return(paginationOut, nil)

	permService := service.NewPermissionService(r)
	permRes, err := permService.FindAll(test.Context{}, &proto.FindAllPermissionRequest{Limit: 10, Page: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes, fmt.Sprintf("Want %v but got %v", want, permRes))
}

func (t *PermissionServiceTest) TestFindOnePermission() {
	var errs []string

	want := &proto.PermissionResponse{
		Data:       test.RawToDtoPermission(t.Perm),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &permission.MockRepo{}

	r.On("FindOne", 1, &model.Permission{}).Return(t.Perm, nil)

	permService := service.NewPermissionService(r)
	permRes, err := permService.FindOne(test.Context{}, &proto.FindOnePermissionRequest{Id: 1})
	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestFindOneErrNotFoundPermission() {
	errs := []string{"Not found permission"}

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &permission.MockRepo{}

	r.On("FindOne", 1, &model.Permission{}).Return(nil, errors.New("Not found permission"))

	permService := service.NewPermissionService(r)
	permRes, err := permService.FindOne(test.Context{}, &proto.FindOnePermissionRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestCreatePermission() {
	var errs []string

	want := &proto.PermissionResponse{
		Data:       test.RawToDtoPermission(t.Perm),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	permIn := &model.Permission{
		Name: t.Perm.Name,
		Code: t.Perm.Code,
	}

	r := &permission.MockRepo{}

	r.On("Create", permIn).Return(t.Perm, nil)

	permService := service.NewPermissionService(r)
	permRes, err := permService.Create(test.Context{}, t.CreatePermissionReqMock)
	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestCreateDuplicatedPermission() {
	errs := []string{"Duplicated permission"}

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusUnprocessableEntity,
	}

	permIn := &model.Permission{
		Name: t.Perm.Name,
		Code: t.Perm.Code,
	}

	r := &permission.MockRepo{}

	r.On("Create", permIn).Return(nil, errors.New("Duplicated permission"))

	permService := service.NewPermissionService(r)
	permRes, err := permService.Create(test.Context{}, t.CreatePermissionReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestUpdatePermission() {
	var errs []string

	want := &proto.PermissionResponse{
		Data:       test.RawToDtoPermission(t.Perm),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &permission.MockRepo{}

	r.On("Update", 1, t.Perm).Return(t.Perm, nil)

	permService := service.NewPermissionService(r)
	permRes, err := permService.Update(test.Context{}, t.UpdatePermissionReqMock)
	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestUpdateErrNotFoundPermission() {
	errs := []string{"Not found permission"}

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &permission.MockRepo{}

	r.On("Update", 1, t.Perm).Return(nil, errors.New("Not found permission"))

	permService := service.NewPermissionService(r)
	permRes, err := permService.Update(test.Context{}, t.UpdatePermissionReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestDeletePermission() {
	var errs []string

	want := &proto.PermissionResponse{
		Data:       test.RawToDtoPermission(t.Perm),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &permission.MockRepo{}

	r.On("Delete", 1, &model.Permission{}).Return(t.Perm, nil)

	permService := service.NewPermissionService(r)
	permRes, err := permService.Delete(test.Context{}, &proto.DeletePermissionRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}

func (t *PermissionServiceTest) TestDeleteErrNotFoundPermission() {
	errs := []string{"Not found permission"}

	want := &proto.PermissionResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &permission.MockRepo{}

	r.On("Delete", 1, &model.Permission{}).Return(nil, errors.New("Not found permission"))

	permService := service.NewPermissionService(r)
	permRes, err := permService.Delete(test.Context{}, &proto.DeletePermissionRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, permRes)
}
