package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (s Seed) RoleSeed1651703066048() error {
	permission := model.Role{Name: faker.Name(), Description: faker.Sentence()}

	return s.db.Create(&permission).Error
}
