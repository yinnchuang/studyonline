package entity

import "gorm.io/gorm"

type LessonPlanStudent struct {
	gorm.Model
	Title             string `json:"title"`
	Duration          string `json:"duration"`
	Objectives        string `json:"objectives"`
	KeyPoints         string `json:"key_points"`
	DifficultPoints   string `json:"difficult_points"`
	Content           string `json:"content"`
	IdeologicalPoints string `json:"ideological_points"`
	UnitIds           string `json:"unit_ids"`
	FatherId          uint   `json:"father_id"`
}
