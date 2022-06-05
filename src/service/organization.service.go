package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type OrganizationService struct {
	repository OrganizationRepository
}

type OrganizationRepository interface {
	FindAll(*model.OrganizationPagination) error
	FindOne(uint, *model.Organization) error
	FindMulti([]uint32, *[]*model.Organization) error
	Create(*model.Organization) error
	Update(uint, *model.Organization) error
	Delete(uint, *model.Organization) error
}

func NewOrganizationService(repository OrganizationRepository) *OrganizationService {
	return &OrganizationService{repository: repository}
}

func (s *OrganizationService) FindAll(_ context.Context, req *proto.FindAllOrganizationRequest) (res *proto.OrganizationPaginationResponse, err error) {
	var teams []*model.Organization
	var errors []string

	query := model.OrganizationPagination{
		Items: &teams,
		Meta: model.PaginationMetadata{
			ItemsPerPage: req.Limit,
			CurrentPage:  req.Page,
		},
	}

	res = &proto.OrganizationPaginationResponse{
		Data: &proto.OrganizationPagination{
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

	var result []*proto.Organization

	for _, perm := range *query.Items {
		result = append(result, RawToDtoOrganization(perm))
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

func (s *OrganizationService) FindOne(_ context.Context, req *proto.FindOneOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := model.Organization{}
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoOrganization(&org)
	res.Data = result
	return
}

func (s *OrganizationService) FindMulti(_ context.Context, req *proto.FindMultiOrganizationRequest) (res *proto.OrganizationListResponse, err error) {
	var orgs []*model.Organization
	var errors []string

	res = &proto.OrganizationListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &orgs)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	var result []*proto.Organization
	for _, org := range orgs {
		result = append(result, RawToDtoOrganization(org))
	}

	res.Data = result

	return
}

func (s *OrganizationService) Create(_ context.Context, req *proto.CreateOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := DtoToRawOrganization(req.Organization)
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoOrganization(org)
	res.Data = result

	return
}

func (s *OrganizationService) Update(_ context.Context, req *proto.UpdateOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := DtoToRawOrganization(req.Organization)
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(uint(org.ID), org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoOrganization(org)
	res.Data = result

	return
}

func (s *OrganizationService) Delete(_ context.Context, req *proto.DeleteOrganizationRequest) (res *proto.OrganizationResponse, err error) {
	org := model.Organization{}
	var errors []string

	res = &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &org)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoOrganization(&org)
	res.Data = result

	return
}
