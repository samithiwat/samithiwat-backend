package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type PermissionService struct {
	repository PermissionRepository
}

type PermissionRepository interface {
	FindAll(*model.PermissionPagination) error
	FindOne(int, *model.Permission) error
	Create(*model.Permission) error
	Update(int, *model.Permission) error
	Delete(int, *model.Permission) error
}

func NewPermissionService(repository PermissionRepository) *PermissionService {
	return &PermissionService{repository: repository}
}

func (s *PermissionService) FindAll(_ context.Context, req *proto.FindAllPermissionRequest) (res *proto.PermissionListResponse, err error) {
	var perms []*model.Permission
	var errors []string

	query := model.PermissionPagination{
		Items: &perms,
		Meta: model.PaginationMetadata{
			ItemsPerPage: req.Limit,
			CurrentPage:  req.Page,
		},
	}

	res = &proto.PermissionListResponse{
		Data: &proto.PermissionPagination{
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

	var result []*proto.Permission

	for _, perm := range *query.Items {
		result = append(result, RawToDtoPermission(perm))
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

func (s *PermissionService) FindOne(_ context.Context, req *proto.FindOnePermissionRequest) (res *proto.PermissionResponse, err error) {
	perm := model.Permission{}
	var errors []string

	res = &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(int(req.Id), &perm)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoPermission(&perm)
	res.Data = result
	return
}

func (s *PermissionService) Create(_ context.Context, req *proto.CreatePermissionRequest) (res *proto.PermissionResponse, err error) {
	perm := DtoToRawPermission(req.Permission)
	var errors []string

	res = &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(perm)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoPermission(perm)
	res.Data = result

	return
}

func (s *PermissionService) Update(_ context.Context, req *proto.UpdatePermissionRequest) (res *proto.PermissionResponse, err error) {
	perm := DtoToRawPermission(req.Permission)
	var errors []string

	res = &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(int(perm.ID), perm)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoPermission(perm)
	res.Data = result

	return
}

func (s *PermissionService) Delete(_ context.Context, req *proto.DeletePermissionRequest) (res *proto.PermissionResponse, err error) {
	perm := model.Permission{}
	var errors []string

	res = &proto.PermissionResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(int(req.Id), &perm)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoPermission(&perm)
	res.Data = result

	return
}
