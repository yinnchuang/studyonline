package entity

type LessonPlan struct {
	ID                uint   `gorm:"primarykey"`
	Title             string `json:"title"`
	Duration          string `json:"duration"`
	Objectives        string `json:"objectives"`
	KeyPoints         string `json:"key_points"`
	DifficultPoints   string `json:"difficult_points"`
	Content           string `json:"content"`
	IdeologicalPoints string `json:"ideological_points"`
}
