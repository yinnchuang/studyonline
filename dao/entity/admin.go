package entity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name      string `json:"name" gorm:"not null"`
	StudentId string `json:"student_id" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
}
