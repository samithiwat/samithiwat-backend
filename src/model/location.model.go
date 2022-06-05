package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Address  string `json:"address"`
	District string `json:"district"`
	Province string `json:"province"`
	Country  string `json:"country"`
	ZipCode  string `json:"zipcode"`
}
