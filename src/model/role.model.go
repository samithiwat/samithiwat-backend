package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Permissions    []*Permission `json:"permissions" gorm:"many2many:role_permission"`
	Users          []*User       `json:"users" gorm:"many2many:user_role;"`
	OrganizationID *uint         `json:"organization_id"`
}

type RolePagination struct {
	Items *[]*Role
	Meta  PaginationMetadata
}
