package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

type ScoreResultVO struct {
	Name       string `json:"name"`
	StudentID  uint   `json:"student_id"`
	UsualScore int    `json:"usual_score"`
	ExamScore  int    `json:"exam_score"`
	FinalScore int    `json:"final_score"`
}

func GetScoreByStudentId(c context.Context, studentId uint) ([]ScoreResultVO, error) {
	var results []ScoreResultVO
	err := mysql.DB.Model(&entity.Score{}).
		Select(`
		students.name,
        scores.student_id,
        scores.usual_score,
		scores.exam_score,
		scores.final_score
    `).
		Joins("JOIN students ON scores.student_id = students.id").
		Where("scores.student_id = ?", studentId).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetAllScore(c context.Context) ([]ScoreResultVO, error) {
	var results []ScoreResultVO
	err := mysql.DB.Model(&entity.Score{}).
		Select(`
		students.name,
        scores.student_id,
        scores.usual_score,
		scores.exam_score,
		scores.final_score
    `).
		Joins("JOIN students ON scores.student_id = students.id").
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func CreateScore(c context.Context, score entity.Score) error {
	if err := mysql.DB.Model(&entity.Score{}).Create(&score).Error; err != nil {
		return err
	}
	return nil
}
