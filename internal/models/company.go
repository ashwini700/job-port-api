package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	// Cid      uint   `json:"cid" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
}
