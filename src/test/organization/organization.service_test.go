package organization_test

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"github.com/samithiwat/samithiwat-backend/src/proto"
	"github.com/samithiwat/samithiwat-backend/src/service"
	"github.com/samithiwat/samithiwat-backend/src/test"
	"github.com/samithiwat/samithiwat-backend/src/test/organization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

type OrganizationServiceTest struct {
	suite.Suite
	Organization              *model.Organization
	Organizations             []*model.Organization
	CreateOrganizationReqMock *proto.CreateOrganizationRequest
	UpdateOrganizationReqMock *proto.UpdateOrganizationRequest
}

func TestOrganizationService(t *testing.T) {
	suite.Run(t, new(OrganizationServiceTest))
}

func (t *OrganizationServiceTest) SetupTest() {
	t.Organization = &model.Organization{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization2 := &model.Organization{
		Model: gorm.Model{
			ID:        2,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization3 := &model.Organization{
		Model: gorm.Model{
			ID:        3,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization4 := &model.Organization{
		Model: gorm.Model{
			ID:        4,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	t.Organizations = append(t.Organizations, t.Organization, Organization2, Organization3, Organization4)

	t.CreateOrganizationReqMock = &proto.CreateOrganizationRequest{
		Organization: &proto.Organization{
			Name:        t.Organization.Name,
			Description: t.Organization.Description,
		},
	}

	t.UpdateOrganizationReqMock = &proto.UpdateOrganizationRequest{
		Organization: &proto.Organization{
			Id:          uint32(t.Organization.ID),
			Name:        t.Organization.Name,
			Description: t.Organization.Description,
		},
	}
}

func (t *OrganizationServiceTest) TestFindAllOrganization() {

	var result []*proto.Organization
	for _, org := range t.Organizations {
		result = append(result, test.RawToDtoOrganization(org))
	}

	var errs []string

	want := &proto.OrganizationPaginationResponse{
		Data: &proto.OrganizationPagination{
			Items: result,
			Meta: &proto.PaginationMetadata{
				ItemsPerPage: 10,
				ItemCount:    int64(len(t.Organizations)),
				TotalItem:    4,
				CurrentPage:  1,
				TotalPage:    1,
			},
		},
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var teams []*model.Organization

	paginationIn := &model.OrganizationPagination{
		Items: &teams,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			CurrentPage:  1,
		},
	}

	paginationOut := model.OrganizationPagination{
		Items: &t.Organizations,
		Meta: model.PaginationMetadata{
			ItemsPerPage: 10,
			ItemCount:    int64(len(t.Organizations)),
			TotalItem:    4,
			CurrentPage:  1,
			TotalPage:    1,
		},
	}

	r := &organization.MockRepo{}

	r.On("FindAll", paginationIn).Return(paginationOut, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.FindAll(test.Context{}, &proto.FindAllOrganizationRequest{Limit: 10, Page: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes, fmt.Sprintf("Want %v but got %v", want, orgRes))
}

func (t *OrganizationServiceTest) TestFindOneOrganization() {
	var errs []string

	want := &proto.OrganizationResponse{
		Data:       test.RawToDtoOrganization(t.Organization),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &organization.MockRepo{}

	r.On("FindOne", uint(1), &model.Organization{}).Return(t.Organization, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.FindOne(test.Context{}, &proto.FindOneOrganizationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestFindOneErrNotFoundOrganization() {
	errs := []string{"Not found organization"}

	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &organization.MockRepo{}

	r.On("FindOne", uint(1), &model.Organization{}).Return(nil, errors.New("Not found organization"))

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.FindOne(test.Context{}, &proto.FindOneOrganizationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestFindMultiOrganization() {
	var result []*proto.Organization
	for _, org := range t.Organizations {
		result = append(result, test.RawToDtoOrganization(org))
	}

	var errs []string

	want := &proto.OrganizationListResponse{
		Data:       result,
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	var orgs []*model.Organization

	r := &organization.MockRepo{}

	r.On("FindMulti", []uint32{1, 2, 3, 4, 5}, &orgs).Return(t.Organizations, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.FindMulti(test.Context{}, &proto.FindMultiOrganizationRequest{Ids: []uint32{1, 2, 3, 4, 5}})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestCreateOrganization() {
	var errs []string

	want := &proto.OrganizationResponse{
		Data:       test.RawToDtoOrganization(t.Organization),
		Errors:     errs,
		StatusCode: http.StatusCreated,
	}

	orgIn := &model.Organization{
		Name:        t.Organization.Name,
		Description: t.Organization.Description,
	}

	r := &organization.MockRepo{}

	r.On("Create", orgIn).Return(t.Organization, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.Create(test.Context{}, t.CreateOrganizationReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestUpdateOrganization() {
	var errs []string

	want := &proto.OrganizationResponse{
		Data:       test.RawToDtoOrganization(t.Organization),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &organization.MockRepo{}

	r.On("Update", uint(1), t.Organization).Return(t.Organization, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.Update(test.Context{}, t.UpdateOrganizationReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestUpdateErrNotFoundOrganization() {
	errs := []string{"Not found organization"}

	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &organization.MockRepo{}

	r.On("Update", uint(1), t.Organization).Return(nil, errors.New("Not found organization"))

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.Update(test.Context{}, t.UpdateOrganizationReqMock)

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestDeleteOrganization() {
	var errs []string

	want := &proto.OrganizationResponse{
		Data:       test.RawToDtoOrganization(t.Organization),
		Errors:     errs,
		StatusCode: http.StatusOK,
	}

	r := &organization.MockRepo{}

	r.On("Delete", uint(1), &model.Organization{}).Return(t.Organization, nil)

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.Delete(test.Context{}, &proto.DeleteOrganizationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}

func (t *OrganizationServiceTest) TestDeleteErrNotFoundOrganization() {
	errs := []string{"Not found organization"}

	want := &proto.OrganizationResponse{
		Data:       nil,
		Errors:     errs,
		StatusCode: http.StatusNotFound,
	}

	r := &organization.MockRepo{}

	r.On("Delete", uint(1), &model.Organization{}).Return(nil, errors.New("Not found organization"))

	orgService := service.NewOrganizationService(r)
	orgRes, err := orgService.Delete(test.Context{}, &proto.DeleteOrganizationRequest{Id: 1})

	assert.Nil(t.T(), err, "Must not got error")
	assert.Equal(t.T(), want, orgRes)
}
