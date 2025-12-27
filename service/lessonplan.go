package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetAllLessonPlan() ([]entity.LessonPlan, error) {
	var res []entity.LessonPlan
	err := mysql.DB.Model(&entity.LessonPlan{}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func RemoveLessonPlan(ctx context.Context, id uint) error {
	err := mysql.DB.Where("id = ?", id).Delete(&entity.LessonPlan{}).Error
	if err != nil {
		return err
	}
	return nil
}
