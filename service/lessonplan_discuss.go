package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"github.com/gin-gonic/gin"
)

func GetLessonPlanDiscussByID(c *gin.Context, id uint) (*entity.LessonPlanDiscuss, error) {
	var discuss entity.LessonPlanDiscuss
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Where("id = ?", id).Find(&discuss).Error
	if err != nil {
		return nil, err
	}
	return &discuss, nil
}

func GetLessonPlanDiscussByLessonPlanID(c *gin.Context, id uint) (*entity.LessonPlanDiscuss, error) {
	var discuss entity.LessonPlanDiscuss
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Where("lesson_plan_id = ?", id).Find(&discuss).Error
	if err != nil {
		return nil, err
	}
	return &discuss, nil
}

func CreateLessonPlanDiscuss(c context.Context, discuss entity.LessonPlanDiscuss) error {
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Create(&discuss).Error
	return err
}

func RemoveLessonPlanDiscussByID(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.LessonPlanDiscuss{}).Delete(&entity.LessonPlanDiscuss{}, id).Error
	return err
}

func RemoveLessonPlanDiscussByFatherID(c context.Context, id uint) error {
	err := mysql.DB.Where("father_id = ?", id).Delete(&entity.LessonPlanDiscuss{}).Error
	return err
}
