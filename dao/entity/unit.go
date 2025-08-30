package entity

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	UnitName string `json:"unit_name" gorm:"not null"`
	UnitDesc string `json:"unit_desc"`
}
