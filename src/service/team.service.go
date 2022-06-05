package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type TeamRepository interface {
	FindAll(*model.TeamPagination) error
	FindOne(uint, *model.Team) error
	FindMulti([]uint32, *[]*model.Team) error
	Create(*model.Team) error
	Update(uint, *model.Team) error
	Delete(uint, *model.Team) error
}

type TeamService struct {
	repository TeamRepository
}

func NewTeamService(repository TeamRepository) *TeamService {
	return &TeamService{
		repository: repository,
	}
}

func (s *TeamService) FindAll(_ context.Context, req *proto.FindAllTeamRequest) (res *proto.TeamPaginationResponse, err error) {
	var teams []*model.Team
	var errors []string

	query := model.TeamPagination{
		Items: &teams,
		Meta: model.PaginationMetadata{
			ItemsPerPage: req.Limit,
			CurrentPage:  req.Page,
		},
	}

	res = &proto.TeamPaginationResponse{
		Data: &proto.TeamPagination{
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

	var result []*proto.Team

	for _, perm := range *query.Items {
		result = append(result, RawToDtoTeam(perm))
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

func (s *TeamService) FindMulti(_ context.Context, req *proto.FindMultiTeamRequest) (res *proto.TeamListResponse, err error) {
	var teams []*model.Team
	var errors []string

	res = &proto.TeamListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &teams)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	var result []*proto.Team
	for _, team := range teams {
		result = append(result, RawToDtoTeam(team))
	}

	res.Data = result

	return
}

func (s *TeamService) FindOne(_ context.Context, req *proto.FindOneTeamRequest) (res *proto.TeamResponse, err error) {
	t := model.Team{}
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(uint(req.Id), &t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoTeam(&t)
	res.Data = result

	return
}

func (s *TeamService) Create(_ context.Context, req *proto.CreateTeamRequest) (res *proto.TeamResponse, err error) {
	t := DtoToRawTeam(req.Team)
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoTeam(t)
	res.Data = result

	return
}

func (s *TeamService) Update(_ context.Context, req *proto.UpdateTeamRequest) (res *proto.TeamResponse, err error) {
	t := DtoToRawTeam(req.Team)
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(uint(t.ID), t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoTeam(t)
	res.Data = result

	return
}

func (s *TeamService) Delete(_ context.Context, req *proto.DeleteTeamRequest) (res *proto.TeamResponse, err error) {
	t := model.Team{}
	var errors []string

	res = &proto.TeamResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(uint(req.Id), &t)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoTeam(&t)
	res.Data = result

	return
}
