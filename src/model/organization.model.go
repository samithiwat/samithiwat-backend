package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string   `json:"name" gorm:"index:,unique"`
	Description string   `json:"description"`
	Teams       []*Team  `json:"teams"`
	Members     []*User  `json:"users" gorm:"many2many:user_organization;"`
	Roles       []*Role  `json:"roles"`
	Contact     Contact  `json:"contact" gorm:"polymorphic:Owner;"`
	Location    Location `json:"location" gorm:"polymorphic:Owner;"`
}

type OrganizationPagination struct {
	Items *[]*Organization
	Meta  PaginationMetadata
}
