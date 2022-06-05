package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{
		db: db,
	}
}

func (r *LocationRepository) FindOne(id int, loc *model.Location) error {
	return r.db.Preload(clause.Associations).First(&loc, id).Error
}

func (r *LocationRepository) FindMulti(ids []uint32, locations *[]*model.Location) error {
	return r.db.Where("id IN ?", ids).Find(&locations).Error
}

func (r *LocationRepository) Create(loc *model.Location) error {
	return r.db.Create(&loc).Error
}

func (r *LocationRepository) Update(id int, loc *model.Location) error {
	return r.db.Where(id).Updates(&loc).First(&loc).Error
}

func (r *LocationRepository) Delete(id int, loc *model.Location) error {
	return r.db.First(&loc, id).Delete(&model.Location{}).Error
}
