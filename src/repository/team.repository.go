package repository

import (
	"github.com/samithiwat/samithiwat-backend/src/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (r *TeamRepository) FindAll(pagination *model.TeamPagination) error {
	return r.db.Scopes(Pagination(&pagination.Items, &pagination.Meta, r.db)).Find(&pagination.Items).Error
}

func (r *TeamRepository) FindOne(id uint, team *model.Team) error {
	return r.db.Preload(clause.Associations).First(&team, id).Error
}

func (r *TeamRepository) FindMulti(ids []uint32, teams *[]*model.Team) error {
	return r.db.Where("id IN ?", ids).Find(&teams).Error
}

func (r *TeamRepository) Create(team *model.Team) error {
	return r.db.Create(&team).Error
}

func (r *TeamRepository) Update(id uint, team *model.Team) error {
	return r.db.Where(id).Updates(&team).First(&team).Error
}

func (r *TeamRepository) Delete(id uint, team *model.Team) error {
	return r.db.First(&team, id).Delete(&model.Team{}).Error
}
