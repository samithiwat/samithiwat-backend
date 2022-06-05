package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	Linkedin  string `json:"linkedin"`
}
