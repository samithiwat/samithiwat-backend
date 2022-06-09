package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type RoleService struct {
	repository RoleRepository
}

type RoleRepository interface {
	FindAll(*model.RolePagination) error
	FindOne(int, *model.Role) error
	FindMulti([]uint32, *[]*model.Role) error
	Create(*model.Role) error
	Update(int, *model.Role) error
	Delete(int, *model.Role) error
}

func NewRoleService(repository RoleRepository) *RoleService {
	return &RoleService{repository: repository}
}

func (s *RoleService) FindAll(_ context.Context, req *proto.FindAllRoleRequest) (res *proto.RolePaginationResponse, err error) {
	var roles []*model.Role
	var errors []string

	query := model.RolePagination{
		Items: &roles,
		Meta: model.PaginationMetadata{
			ItemsPerPage: req.Limit,
			CurrentPage:  req.Page,
		},
	}

	res = &proto.RolePaginationResponse{
		Data: &proto.RolePagination{
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

	var result []*proto.Role

	for _, role := range *query.Items {
		result = append(result, RawToDtoRole(role))
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

func (s *RoleService) FindOne(_ context.Context, req *proto.FindOneRoleRequest) (res *proto.RoleResponse, err error) {
	role := model.Role{}
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(int(req.Id), &role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoRole(&role)
	res.Data = result

	return
}

func (s *RoleService) FindMulti(_ context.Context, req *proto.FindMultiRoleRequest) (res *proto.RoleListResponse, err error) {
	var roles []*model.Role
	var errors []string

	res = &proto.RoleListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &roles)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	var result []*proto.Role
	for _, role := range roles {
		result = append(result, RawToDtoRole(role))
	}

	res.Data = result

	return
}

func (s *RoleService) Create(_ context.Context, req *proto.CreateRoleRequest) (res *proto.RoleResponse, err error) {
	role := DtoToRawRole(req.Role)
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoRole(role)
	res.Data = result

	return
}

func (s *RoleService) Update(_ context.Context, req *proto.UpdateRoleRequest) (res *proto.RoleResponse, err error) {
	role := DtoToRawRole(req.Role)
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(int(role.ID), role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoRole(role)
	res.Data = result

	return
}

func (s *RoleService) Delete(_ context.Context, req *proto.DeleteRoleRequest) (res *proto.RoleResponse, err error) {
	role := model.Role{}
	var errors []string

	res = &proto.RoleResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(int(req.Id), &role)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoRole(&role)
	res.Data = result

	return
}
