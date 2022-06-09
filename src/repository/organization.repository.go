package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) FindAll(pagination *model.OrganizationPagination) error {
	return r.db.Scopes(Pagination(&pagination.Items, &pagination.Meta, r.db)).Find(&pagination.Items).Error
}

func (r *OrganizationRepository) FindOne(id uint, perm *model.Organization) error {
	return r.db.Preload(clause.Associations).First(&perm, id).Error
}

func (r *OrganizationRepository) FindMulti(ids []uint32, orgs *[]*model.Organization) error {
	return r.db.Where("id IN ?", ids).Find(&orgs).Error
}

func (r *OrganizationRepository) Create(perm *model.Organization) error {
	return r.db.Create(&perm).Error
}

func (r *OrganizationRepository) Update(id uint, perm *model.Organization) error {
	return r.db.Where(id).Updates(&perm).First(&perm).Error
}

func (r *OrganizationRepository) Delete(id uint, perm *model.Organization) error {
	return r.db.First(&perm, id).Delete(&model.Organization{}).Error
}
