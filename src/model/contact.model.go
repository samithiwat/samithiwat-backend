package model

import "gorm.io/gorm"

// TODO: Change to use polymorph relationship

type Contact struct {
	gorm.Model
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
	OwnerID   *uint  `json:"owner_id"`
	OwnerType string `json:"owner_type"`
}
