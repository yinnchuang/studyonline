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

func UpdateScore(c context.Context, score entity.Score) error {
	if err := mysql.DB.Model(&entity.Score{}).Where("student_id = ?", score.StudentId).Updates(&score).Error; err != nil {
		return err
	}
	return nil
}

func ExistScore(c context.Context, studentID uint) (bool, error) {
	var count int64
	if err := mysql.DB.Model(&entity.Score{}).Where("student_id = ?", studentID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
