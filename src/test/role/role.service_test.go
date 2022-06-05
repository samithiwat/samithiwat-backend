package role_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/role"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type RoleServiceTest struct {
	suite.Suite
	Role              *model.Role
	Roles             []*model.Role
	CreateRoleReqMock *proto.CreateRoleRequest
	UpdateRoleReqMock *proto.UpdateRoleRequest
}

func TestRoleService(t *testing.T) {
	suite.Run(t, new(RoleServiceTest))
}

func (t *RoleServiceTest) SetupTest() {
	t.Role = &model.Role{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role2 := &model.Role{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role3 := &model.Role{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Role4 := &model.Role{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	t.Roles = append(t.Roles, t.Role, Role2, Role3, Role4)

	t.CreateRoleReqMock = &proto.CreateRoleRequest{
		Role: &proto.Role{
			Name:        t.Role.Name,
			Description: t.Role.Description,
		},
	}

	t.UpdateRoleReqMock = &proto.UpdateRoleRequest{
		Role: &proto.Role{
			Id:          uint32(t.Role.ID),
			Name:        t.Role.Name,
			Description: t.Role.Description,
		},
	}
}

func (t *RoleServiceTest) TestFindAllRole() {
	var result []*proto.Role
	for _, role := range t.Roles {
		result = append(result, test.RawToDtoRole(role))
	}

	var errs []string

	want := &proto.RolePaginationResponse{
		Data: &proto.RolePagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    int64(len(t.Roles)),
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var perms []*model.Role

	paginationIn := &model.RolePagination{
		Items: &perms,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			CurrentPage:  1,
		},
	}

	paginationOut := model.RolePagination{
		Items: &t.Roles,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			ItemCount:    int64(len(t.Roles)),
			TotalItem:    4,
			CurrentPage:  1,
			TotalPage:    1,
		},
	}

	r := &role.MockRepo{}

	r.On("FindAll", paginationIn).Return(paginationOut, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.FindAll(test.Context{}, &proto.FindAllRoleRequest{Limit: 10, Page: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes, fmt.Sprintf("Want %v but got %v", want, roleRes))
}

func (t *RoleServiceTest) TestFindOneRole() {
	var errs []string
	want := &proto.RoleResponse{
		Data:       test.RawToDtoRole(t.Role),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &role.MockRepo{}

	r.On("FindOne", 1, &model.Role{}).Return(t.Role, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.FindOne(test.Context{}, &proto.FindOneRoleRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestFindOneErrNotFoundRole() {
	errs := []string{"Not found role"}
	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &role.MockRepo{}

	r.On("FindOne", 1, &model.Role{}).Return(nil, errors.New("Not found role"))

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.FindOne(test.Context{}, &proto.FindOneRoleRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestFindMultiRole() {
	var result []*proto.Role
	for _, role := range t.Roles {
		result = append(result, test.RawToDtoRole(role))
	}

	var errs []string
	want := &proto.RoleListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &role.MockRepo{}

	var roles []*model.Role

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &roles).Return(t.Roles, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.FindMulti(test.Context{}, &proto.FindMultiRoleRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestCreateRole() {
	var errs []string
	want := &proto.RoleResponse{
		Data:       test.RawToDtoRole(t.Role),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	roleIn := &model.Role{
		Name:        t.Role.Name,
		Description: t.Role.Description,
	}

	r := &role.MockRepo{}

	r.On("Create", roleIn).Return(t.Role, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.Create(test.Context{}, t.CreateRoleReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestUpdateRole() {
	var errs []string
	want := &proto.RoleResponse{
		Data:       test.RawToDtoRole(t.Role),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &role.MockRepo{}

	r.On("Update", 1, t.Role).Return(t.Role, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.Update(test.Context{}, t.UpdateRoleReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestUpdateErrNotFoundRole() {
	errs := []string{"Not found role"}

	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &role.MockRepo{}

	r.On("Update", 1, t.Role).Return(nil, errors.New("Not found role"))

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.Update(test.Context{}, t.UpdateRoleReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestDeleteRole() {
	var errs []string
	want := &proto.RoleResponse{
		Data:       test.RawToDtoRole(t.Role),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &role.MockRepo{}

	r.On("Delete", 1, &model.Role{}).Return(t.Role, nil)

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.Delete(test.Context{}, &proto.DeleteRoleRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}

func (t *RoleServiceTest) TestDeleteErrNotFoundRole() {
	errs := []string{"Not found role"}

	want := &proto.RoleResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &role.MockRepo{}

	r.On("Delete", 1, &model.Role{}).Return(t.Role, errors.New("Not found role"))

	roleService := service.NewRoleService(r)
	roleRes, err := roleService.Delete(test.Context{}, &proto.DeleteRoleRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, roleRes)
}
