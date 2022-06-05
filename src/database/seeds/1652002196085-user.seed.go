package seed

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend/src/model"
)

func (s Seed) UserSeed1652002196085() error {
	for i := 0; i < 10; i++ {
		user := model.User{Firstname: faker.FirstName(),
			Lastname:    faker.LastName(),
			ImageUrl:    faker.URL(),
			DisplayName: faker.Username(),
		}
		err := s.db.Create(&user).Error

		if err != nil {
			return err
		}
	}
	return nil
}
