package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

// GetAllLessonPlan 教师端教案列表：已发布的全部教案 + 自己创建的草稿
func GetAllLessonPlan(teacherID uint) ([]entity.LessonPlan, error) {
	var res []entity.LessonPlan
	err := mysql.DB.Model(&entity.LessonPlan{}).
		Where("publish_status = 1 OR teacher_id = ?", teacherID).
		Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RemoveLessonPlan 仅允许删除自己创建的教案
func RemoveLessonPlan(ctx context.Context, id uint, teacherID uint) error {
	err := mysql.DB.Where("id = ? AND teacher_id = ?", id, teacherID).Delete(&entity.LessonPlan{}).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateLessonPlan 仅允许修改自己创建的教案
func UpdateLessonPlan(ctx context.Context, id uint, teacherID uint, lp *entity.LessonPlan) error {
	err := mysql.DB.Model(&entity.LessonPlan{}).
		Where("id = ? AND teacher_id = ?", id, teacherID).
		Updates(lp).Error
	if err != nil {
		return err
	}
	return nil
}

// GetLessonPlanById 根据ID获取教案（需校验所有权）
func GetLessonPlanById(ctx context.Context, id uint, teacherID uint) (*entity.LessonPlan, error) {
	var lp *entity.LessonPlan
	err := mysql.DB.Where("id = ? AND teacher_id = ?", id, teacherID).First(&lp).Error
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

func GetLessonPlanStudentByFatherId(ctx context.Context, fatherId uint) (*entity.LessonPlanStudent, error) {
	var lp *entity.LessonPlanStudent
	err := mysql.DB.Where("father_id = ?", fatherId).First(&lp).Error
	if err != nil {
		return nil, err
	}
	return lp, nil
}

func UpdateLessonPlanStudent(ctx context.Context, lp *entity.LessonPlanStudent) error {
	err := mysql.DB.Where("id = ?", lp.ID).Updates(lp).Error
	if err != nil {
		return err
	}
	return nil
}
