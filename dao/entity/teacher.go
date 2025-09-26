package entity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null"`
	Username string `json:"Username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
