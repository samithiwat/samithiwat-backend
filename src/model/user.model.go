package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname     string          `json:"firstname"`
	Lastname      string          `json:"lastname"`
	ImageUrl      string          `json:"image_url"`
	DisplayName   string          `json:"display_name"`
	Organizations []*Organization `json:"organizations" gorm:"many2many:user_organization;"`
	Teams         []*Team         `json:"teams" gorm:"many2many:user_team;"`
	Roles         []*Role         `json:"roles" gorm:"many2many:user_role;"`
	Location      Location        `json:"location" gorm:"polymorphic:Owner;"`
	Contact       Contact         `json:"contact" gorm:"polymorphic:Owner;"`
}

type UserPagination struct {
	Items *[]*User
	Meta  PaginationMetadata
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	var count int64
	err = db.Model(u).Where("display_name LIKE ?", fmt.Sprintf("%%%v%%", u.DisplayName)).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		u.DisplayName = fmt.Sprintf("%v#%v", u.DisplayName, count+1)
	}

	return nil
}
