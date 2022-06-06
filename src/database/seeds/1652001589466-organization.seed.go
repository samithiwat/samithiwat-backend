package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/model"
	"math/rand"
)

func (s Seed) OrganizationSeed1652001589466() error {
	faker.SetGenerateUniqueValues(true)
	org := model.Organization{Name: faker.Word(), Description: faker.Sentence(), Email: faker.Email()}
	err := s.db.Create(&org).Error
	if err != nil {
		return err
	}

	for i := 0; i < rand.Intn(10); i++ {
		team := model.Team{Name: faker.Word(), Description: faker.Sentence(), OrganizationID: &org.ID}
		err := s.db.Create(&team).Error
		if err != nil {
			return err
		}
	}

	return nil
}
