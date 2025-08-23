package service

import (
	"context"
	"errors"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func Import(ctx context.Context, user interface{}, identity int) error {
	if identity == constant.StudentIdentity {
		stu := user.(entity.Student)
		err := mysql.DB.Model(&entity.Student{}).Create(&stu).Error
		if err != nil {
			return err
		}
	} else if identity == constant.TeacherIdentity {
		tea := user.(entity.Teacher)
		err := mysql.DB.Model(&entity.Teacher{}).Create(&tea).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New("导入失败")
	}
	return nil
}
