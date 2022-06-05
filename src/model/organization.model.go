package model

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	Name        string   `json:"name" gorm:"index:,unique"`
	Description string   `json:"description"`
	Teams       []*Team  `json:"teams"`
	Roles       []*Role  `json:"roles" `
	Contact     Contact  `json:"contact"`
	Location    Location `json:"location"`
}

type OrganizationPagination struct {
	Items *[]*Organization
	Meta  PaginationMetadata
}
