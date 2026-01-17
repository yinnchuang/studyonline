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

func UpdateLessonPlan(ctx context.Context, id uint, lp *entity.LessonPlan) error {
	err := mysql.DB.Model(&entity.LessonPlan{}).Where("id = ?", id).Updates(lp).Error
	if err != nil {
		return err
	}
	return nil
}

func GetLessonPlanById(ctx context.Context, id uint) (*entity.LessonPlan, error) {
	var lp *entity.LessonPlan
	err := mysql.DB.Where("id = ?", id).First(&lp).Error
	if err != nil {
		return nil, err
	}
	return lp, nil
}

func GetAllLessonPlanStudent() ([]entity.LessonPlanStudent, error) {
	var res []entity.LessonPlanStudent
	err := mysql.DB.Model(&entity.LessonPlanStudent{}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateLessonPlanStudent(ctx context.Context, lp *entity.LessonPlanStudent) error {
	err := mysql.DB.Model(&entity.LessonPlanStudent{}).Create(&lp).Error
	if err != nil {
		return err
	}
	return nil
}

func RemoveLessonPlanStudent(ctx context.Context, fatherId uint) error {
	err := mysql.DB.Where("father_id = ?", fatherId).Delete(&entity.LessonPlanStudent{}).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateLessonPlanStudent(ctx context.Context, lp *entity.LessonPlanStudent) error {
	err := mysql.DB.Where("id = ?", lp.ID).Updates(lp).Error
	if err != nil {
		return err
	}
	return nil
}
