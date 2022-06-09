package user_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type UserServiceTest struct {
	suite.Suite
	User              *model.User
	Users             []*model.User
	CreateUserReqMock *proto.CreateUserRequest
	UpdateUserReqMock *proto.UpdateUserRequest
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.User = &model.User{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		ImageUrl:    faker.URL(),
		DisplayName: faker.Username(),
	}

	User2 := &model.User{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		ImageUrl:    faker.URL(),
		DisplayName: faker.Username(),
	}

	User3 := &model.User{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		ImageUrl:    faker.URL(),
		DisplayName: faker.Username(),
	}

	User4 := &model.User{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		ImageUrl:    faker.URL(),
		DisplayName: faker.Username(),
	}

	t.Users = append(t.Users, t.User, User2, User3, User4)

	t.CreateUserReqMock = &proto.CreateUserRequest{
		User: &proto.User{
			Firstname:   t.User.Firstname,
			Lastname:    t.User.Lastname,
			DisplayName: t.User.DisplayName,
			ImageUrl:    t.User.ImageUrl,
		},
	}

	t.UpdateUserReqMock = &proto.UpdateUserRequest{
		User: &proto.User{
			Id:          uint32(t.User.ID),
			Firstname:   t.User.Firstname,
			Lastname:    t.User.Lastname,
			DisplayName: t.User.DisplayName,
			ImageUrl:    t.User.ImageUrl,
		},
	}
}

func (t *UserServiceTest) TestFindAllUser() {
	var result []*proto.User
	for _, usr := range t.Users {
		result = append(result, test.RawToDtoUser(usr))
	}

	var errs []string

	want := &proto.UserPaginationResponse{
		Data: &proto.UserPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    int64(len(t.Users)),
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var users []*model.User

	paginationIn := &model.UserPagination{
		Items: &users,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			CurrentPage:  1,
		},
	}

	paginationOut := model.UserPagination{
		Items: &t.Users,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			ItemCount:    int64(len(t.Users)),
			TotalItem:    4,
			CurrentPage:  1,
			TotalPage:    1,
		},
	}

	r := &user.MockRepo{}

	r.On("FindAll", paginationIn).Return(paginationOut, nil)

	usrService := service.NewUserService(r)
	usrRes, err := usrService.FindAll(test.Context{}, &proto.FindAllUserRequest{Limit: 10, Page: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, usrRes, fmt.Sprintf("Want %v but got %v", want, usrRes))
}

func (t *UserServiceTest) TestFindOneUser() {
	var errs []string

	want := &proto.UserResponse{
		Data:       test.RawToDtoUser(t.User),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &user.MockRepo{}

	r.On("FindOne", uint(1), &model.User{}).Return(t.User, nil)

	usrService := service.NewUserService(r)
	usrRes, err := usrService.FindOne(test.Context{}, &proto.FindOneUserRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestFindOneErrNotFoundUser() {
	errs := []string{"Not found user"}

	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &user.MockRepo{}

	r.On("FindOne", uint(1), &model.User{}).Return(nil, errors.New("Not found user"))

	usrService := service.NewUserService(r)
	usrRes, _ := usrService.FindOne(test.Context{}, &proto.FindOneUserRequest{Id: 1})

	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestFindMultiUser() {
	var result []*proto.User
	for _, tm := range t.Users {
		result = append(result, test.RawToDtoUser(tm))
	}

	var errs []string

	want := &proto.UserListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var teams []*model.User

	r := &user.MockRepo{}

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &teams).Return(t.Users, nil)

	teamService := service.NewUserService(r)
	teamRes, err := teamService.FindMulti(test.Context{}, &proto.FindMultiUserRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *UserServiceTest) TestCreateUser() {
	var errs []string

	want := &proto.UserResponse{
		Data:       test.RawToDtoUser(t.User),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	userIn := &model.User{
		Firstname:   t.User.Firstname,
		Lastname:    t.User.Lastname,
		ImageUrl:    t.User.ImageUrl,
		DisplayName: t.User.DisplayName,
	}

	r := &user.MockRepo{}

	r.On("Create", userIn).Return(t.User, nil)

	usrService := service.NewUserService(r)
	usrRes, err := usrService.Create(test.Context{}, t.CreateUserReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestUpdateUser() {
	var errs []string

	want := &proto.UserResponse{
		Data:       test.RawToDtoUser(t.User),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &user.MockRepo{}

	r.On("Update", uint(1), t.User).Return(t.User, nil)

	usrService := service.NewUserService(r)
	usrRes, err := usrService.Update(test.Context{}, t.UpdateUserReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestUpdateErrNotFoundUser() {
	errs := []string{"Not found user"}

	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &user.MockRepo{}

	r.On("Update", uint(1), t.User).Return(nil, errors.New("Not found user"))

	usrService := service.NewUserService(r)
	usrRes, _ := usrService.Update(test.Context{}, t.UpdateUserReqMock)

	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestDeleteUser() {
	var errs []string

	want := &proto.UserResponse{
		Data:       test.RawToDtoUser(t.User),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &user.MockRepo{}

	r.On("Delete", uint(1), &model.User{}).Return(t.User, nil)

	usrService := service.NewUserService(r)
	usrRes, err := usrService.Delete(test.Context{}, &proto.DeleteUserRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, usrRes)
}

func (t *UserServiceTest) TestDeleteErrNotFoundUser() {
	errs := []string{"Not found user"}

	want := &proto.UserResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &user.MockRepo{}

	r.On("Delete", uint(1), &model.User{}).Return(nil, errors.New("Not found user"))

	usrService := service.NewUserService(r)
	usrRes, _ := usrService.Delete(test.Context{}, &proto.DeleteUserRequest{Id: 1})

	assert.Equal(t.T(), want, usrRes)
}
