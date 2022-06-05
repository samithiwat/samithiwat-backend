package team_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/team"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type TeamServiceTest struct {
	suite.Suite
	Team              *model.Team
	Teams             []*model.Team
	CreateTeamReqMock *proto.CreateTeamRequest
	UpdateTeamReqMock *proto.UpdateTeamRequest
}

func TestTeamService(t *testing.T) {
	suite.Run(t, new(TeamServiceTest))
}

func (t *TeamServiceTest) SetupTest() {
	t.Team = &model.Team{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team2 := &model.Team{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team3 := &model.Team{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team4 := &model.Team{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	t.Teams = append(t.Teams, t.Team, Team2, Team3, Team4)

	t.CreateTeamReqMock = &proto.CreateTeamRequest{
		Team: &proto.Team{
			Name:        t.Team.Name,
			Description: t.Team.Description,
		},
	}

	t.UpdateTeamReqMock = &proto.UpdateTeamRequest{
		Team: &proto.Team{
			Id:          uint32(t.Team.ID),
			Name:        t.Team.Name,
			Description: t.Team.Description,
		},
	}
}

func (t *TeamServiceTest) TestFindAllTeam() {
	var result []*proto.Team
	for _, tm := range t.Teams {
		result = append(result, test.RawToDtoTeam(tm))
	}

	var errs []string

	want := &proto.TeamPaginationResponse{
		Data: &proto.TeamPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    int64(len(t.Teams)),
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var teams []*model.Team

	paginationIn := &model.TeamPagination{
		Items: &teams,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			CurrentPage:  1,
		},
	}

	paginationOut := model.TeamPagination{
		Items: &t.Teams,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			ItemCount:    int64(len(t.Teams)),
			TotalItem:    4,
			CurrentPage:  1,
			TotalPage:    1,
		},
	}

	r := &team.MockRepo{}

	r.On("FindAll", paginationIn).Return(paginationOut, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.FindAll(test.Context{}, &proto.FindAllTeamRequest{Limit: 10, Page: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes, fmt.Sprintf("Want %v but got %v", want, teamRes))
}

func (t *TeamServiceTest) TestFindOneTeam() {
	var errs []string

	want := &proto.TeamResponse{
		Data:       test.RawToDtoTeam(t.Team),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &team.MockRepo{}

	r.On("FindOne", uint(1), &model.Team{}).Return(t.Team, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.FindOne(test.Context{}, &proto.FindOneTeamRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestFindOneErrNotFoundTeam() {
	errs := []string{"Not found team"}

	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &team.MockRepo{}

	r.On("FindOne", uint(1), &model.Team{}).Return(nil, errors.New("Not found team"))

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.FindOne(test.Context{}, &proto.FindOneTeamRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestFindMultiTeam() {
	var result []*proto.Team
	for _, tm := range t.Teams {
		result = append(result, test.RawToDtoTeam(tm))
	}

	var errs []string

	want := &proto.TeamListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var teams []*model.Team

	r := &team.MockRepo{}

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &teams).Return(t.Teams, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.FindMulti(test.Context{}, &proto.FindMultiTeamRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestCreateTeam() {
	var errs []string

	want := &proto.TeamResponse{
		Data:       test.RawToDtoTeam(t.Team),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	teamIn := &model.Team{
		Name:        t.Team.Name,
		Description: t.Team.Description,
	}

	r := &team.MockRepo{}

	r.On("Create", teamIn).Return(t.Team, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.Create(test.Context{}, t.CreateTeamReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestUpdateTeam() {
	var errs []string

	want := &proto.TeamResponse{
		Data:       test.RawToDtoTeam(t.Team),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &team.MockRepo{}

	r.On("Update", uint(1), t.Team).Return(t.Team, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.Update(test.Context{}, t.UpdateTeamReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestUpdateErrNotFoundTeam() {
	errs := []string{"Not found team"}

	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &team.MockRepo{}

	r.On("Update", uint(1), t.Team).Return(nil, errors.New("Not found team"))

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.Update(test.Context{}, t.UpdateTeamReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestDeleteTeam() {
	var errs []string

	want := &proto.TeamResponse{
		Data:       test.RawToDtoTeam(t.Team),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &team.MockRepo{}

	r.On("Delete", uint(1), &model.Team{}).Return(t.Team, nil)

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.Delete(test.Context{}, &proto.DeleteTeamRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}

func (t *TeamServiceTest) TestDeleteErrNotFoundTeam() {
	errs := []string{"Not found team"}

	want := &proto.TeamResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &team.MockRepo{}

	r.On("Delete", uint(1), &model.Team{}).Return(nil, errors.New("Not found team"))

	teamService := service.NewTeamService(r)
	teamRes, err := teamService.Delete(test.Context{}, &proto.DeleteTeamRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, teamRes)
}
