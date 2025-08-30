package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetScoreByStudentId(c context.Context, studentId uint) ([]entity.Score, error) {
	var scores []entity.Score
	err := mysql.DB.Model(&entity.Score{}).Select("student_id", "unit_id", "score").Where("student_id = ?", studentId).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func GetAllScore(c context.Context) ([]entity.Score, error) {
	var scores []entity.Score
	err := mysql.DB.Model(&entity.Score{}).Select("student_id", "unit_id", "score").Order("student_id").Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func CreateScore(c context.Context, teacherId uint, studentId uint, unitId uint, score int) error {
	cScore := entity.Score{
		StudentId: studentId,
		UnitId:    unitId,
		Score:     score,
		TeacherId: teacherId,
	}
	if err := mysql.DB.Model(&entity.Score{}).Create(&cScore).Error; err != nil {
		return err
	}
	return nil
}
