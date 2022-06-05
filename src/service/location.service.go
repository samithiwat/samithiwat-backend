package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"net/http"
)

type LocationService struct {
	repository LocationRepository
}

type LocationRepository interface {
	FindOne(int, *model.Location) error
	FindMulti([]uint32, *[]*model.Location) error
	Create(*model.Location) error
	Update(int, *model.Location) error
	Delete(int, *model.Location) error
}

func NewLocationService(repository LocationRepository) *LocationService {
	return &LocationService{
		repository: repository,
	}
}

func (s *LocationService) FindOne(_ context.Context, req *proto.FindOneLocationRequest) (res *proto.LocationResponse, err error) {
	loc := model.Location{}
	var errors []string

	res = &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindOne(int(req.Id), &loc)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoLocation(&loc)
	res.Data = result
	return
}

func (s *LocationService) FindMulti(_ context.Context, req *proto.FindMultiLocationRequest) (res *proto.LocationListResponse, err error) {
	var locs []*model.Location
	var errors []string

	res = &proto.LocationListResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.FindMulti(req.Ids, &locs)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	var result []*proto.Location
	for _, loc := range locs {
		result = append(result, RawToDtoLocation(loc))
	}

	res.Data = result

	return
}

func (s *LocationService) Create(_ context.Context, req *proto.CreateLocationRequest) (res *proto.LocationResponse, err error) {
	loc := DtoToRawLocation(req.Location)
	var errors []string

	res = &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusCreated,
	}

	err = s.repository.Create(loc)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	result := RawToDtoLocation(loc)
	res.Data = result

	return
}

func (s *LocationService) Update(_ context.Context, req *proto.UpdateLocationRequest) (res *proto.LocationResponse, err error) {
	loc := DtoToRawLocation(req.Location)
	var errors []string

	res = &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Update(int(loc.ID), loc)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoLocation(loc)
	res.Data = result

	return
}

func (s *LocationService) Delete(_ context.Context, req *proto.DeleteLocationRequest) (res *proto.LocationResponse, err error) {
	loc := model.Location{}
	var errors []string

	res = &proto.LocationResponse{
		Data:       nil,
		Errors:     errors,
		StatusCode: http.StatusOK,
	}

	err = s.repository.Delete(int(req.Id), &loc)
	if err != nil {
		res.Errors = append(errors, err.Error())
		res.StatusCode = http.StatusNotFound
		return res, nil
	}

	result := RawToDtoLocation(&loc)
	res.Data = result

	return
}
