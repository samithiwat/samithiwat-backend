package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	ParentID       *uint   `json:"parent_id"`
	SubTeams       []*Team `json:"sub_team" gorm:"foreignkey:ParentID"`
	Members        []*User `json:"users" gorm:"many2many:user_team;"`
	OrganizationID *uint   `json:"organization_id"`
}

type TeamPagination struct {
	Items *[]*Team
	Meta  PaginationMetadata
}
