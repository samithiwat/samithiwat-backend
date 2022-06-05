package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type UserService struct {
	repository UserRepository
}

type UserRepository interface {
	FindAll(pagination *model.UserPagination) error
	FindOne(uint, *model.User) error
	FindMulti([]uint32, *[]*model.User) error
	Create(*model.User) error
	Update(uint, *model.User) error
	Delete(uint, *model.User) error
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) FindAll(_ context.Context, req *proto.FindAllUserRequest) (res *proto.UserPaginationResponse, err error) {
	var users []*model.User
	var errors []string

	query := model.UserPagination{
		Items: &users,
		Meta: model.PaginationMetadata{
			ItemsPerPage: req.Limit,
			CurrentPage:  req.Page,
		},
	}

	res = &proto.UserPaginationResponse{
		Data: &proto.UserPagination{
			Items: nil,
			Meta:  nil,
		},
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindAll(&query)
	if err != nil {
		errors = append(errors, err.Error())
		res.StatusCode = http.StatusBadRequest
		return res, nil
	}

	var result []*proto.User

	for _, user := range *query.Items {
		result = append(result, RawToDtoUser(user))
	}

	res.Data.Items = result
	res.Data.Meta = &proto.PaginationMetadata{
		TotalItem:    query.Meta.TotalItem,
		ItemCount:    int64(len(result)),
		ItemsPerPage: query.Meta.ItemsPerPage,
		TotalPage:    query.Meta.TotalPage,
		CurrentPage:  query.Meta.CurrentPage,
	}

	return
}

func (s *UserService) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.UserResponse, err error) {
	user := model.User{}
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &user)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoUser(&user)
	res.Data = result
	return
}

func (s *UserService) FindMulti(_ context.Context, req *proto.FindMultiUserRequest) (res *proto.UserListResponse, err error) {
	var users []*model.User
	var errors []string

	res = &proto.UserListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &users)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	var result []*proto.User
	for _, user := range users {
		result = append(result, RawToDtoUser(user))
	}

	res.Data = result

	return
}

func (s *UserService) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.UserResponse, err error) {
	user := DtoToRawUser(req.User)
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(user)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoUser(user)
	res.Data = result

	return
}

func (s *UserService) Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UserResponse, err error) {
	user := DtoToRawUser(req.User)
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(user.ID, user)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoUser(user)
	res.Data = result

	return
}

func (s *UserService) Delete(_ context.Context, req *proto.DeleteUserRequest) (res *proto.UserResponse, err error) {
	user := model.User{}
	var errors []string

	res = &proto.UserResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &user)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoUser(&user)
	res.Data = result

	return
}
