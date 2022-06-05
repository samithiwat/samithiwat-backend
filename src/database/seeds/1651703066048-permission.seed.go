package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (s Seed) PermissionSeed1651703066048() error {
	permission := model.Permission{Name: faker.Word(), Code: faker.Word()}

	return s.db.Create(&permission).Error
}
