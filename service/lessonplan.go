package service

import (
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
