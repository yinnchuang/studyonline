package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name       string `json:"name" gorm:"not null"`
	Username   string `json:"username" gorm:"not null;unique"`
	Password   string `json:"password" gorm:"not null"`
	Department string `json:"department" gorm:"not null"`
}
