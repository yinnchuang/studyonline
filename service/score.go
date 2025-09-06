package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

type ScoreResultVO struct {
	Name          string `json:"name"`
	StudentID     uint   `json:"student_id"`
	HomeworkTitle string `json:"homework_title"`
	HomeworkId    uint   `json:"homework_id"`
	Score         int    `json:"score"`
}

func GetScoreByStudentId(c context.Context, studentId uint) ([]ScoreResultVO, error) {
	var results []ScoreResultVO
	err := mysql.DB.Model(&entity.Score{}).
		Select(`
		students.name,
        scores.student_id, 
        scores.homework_id, 
        scores.score, 
        homeworks.title as homework_title,
    `).
		Joins("JOIN homeworks ON scores.homework_id = homeworks.id").
		Joins("JOIN students ON scores.student_id = students.id").
		Order("scores.homework_id").
		Where("scores.student_id = ?", studentId).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func GetScoreByHomeworkId(c context.Context, homeworkId uint) ([]ScoreResultVO, error) {
	var results []ScoreResultVO
	err := mysql.DB.Model(&entity.Score{}).
		Select(`
		students.name,
        scores.student_id, 
        scores.homework_id, 
        scores.score, 
        homeworks.title as homework_title,
    `).
		Joins("JOIN homeworks ON scores.homework_id = homeworks.id").
		Joins("JOIN students ON scores.student_id = students.id").
		Order("scores.homework_id").
		Where("scores.homework_id = ?", homeworkId).
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
        scores.homework_id, 
        scores.score, 
        homeworks.title as homework_title,
    `).
		Joins("JOIN homeworks ON scores.homework_id = homeworks.id").
		Joins("JOIN students ON scores.student_id = students.id").
		Order("scores.homework_id").
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
