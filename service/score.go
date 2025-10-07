package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetScoreByStudentId(c context.Context, studentId uint) (*entity.Score, error) {
	var score entity.Score
	err := mysql.DB.Model(&entity.Score{}).Where("student_id = ?", studentId).First(&score).Error
	if err != nil {
		return nil, err
	}
	return &score, nil
}

func GetAllScore(c context.Context) ([]entity.Score, error) {
	var scores []entity.Score
	err := mysql.DB.Model(&entity.Score{}).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func CreateScore(c context.Context, score entity.Score) error {
	if err := mysql.DB.Model(&entity.Score{}).Create(&score).Error; err != nil {
		return err
	}
	return nil
}
