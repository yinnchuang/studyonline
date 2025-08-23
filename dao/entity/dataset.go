package entity

import "gorm.io/gorm"

type Dataset struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description,omitempty"`
	FilePath    string `json:"file_path"`
	CoverPath   string `json:"cover_path"`
}
