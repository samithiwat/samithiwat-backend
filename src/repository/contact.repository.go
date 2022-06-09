package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{
		db: db,
	}
}

func (r *ContactRepository) FindOne(id int, cont *model.Contact) error {
	return r.db.Preload(clause.Associations).First(&cont, id).Error
}

func (r *ContactRepository) FindMulti(ids []uint32, conts *[]*model.Contact) error {
	return r.db.Where("id IN ?", ids).Find(&conts).Error
}

func (r *ContactRepository) Create(cont *model.Contact) error {
	return r.db.Create(&cont).Error
}

func (r *ContactRepository) Update(id int, cont *model.Contact) error {
	return r.db.Where(id).Updates(&cont).First(&cont).Error
}

func (r *ContactRepository) Delete(id int, cont *model.Contact) error {
	return r.db.First(&cont, id).Delete(&model.Contact{}).Error
}
