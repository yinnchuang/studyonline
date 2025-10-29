package service

import (
	"context"
	"errors"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func ImportAdmin(ctx context.Context, username string, password string) error {
	result := mysql.DB.Model(&entity.Admin{}).First(&entity.Admin{})
	// 如果找到了记录，则返回
	if result.RowsAffected > 0 {
		return nil
	}
	err := mysql.DB.Model(&entity.Admin{}).Create(&entity.Admin{
		Username: username,
		Password: password,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

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

func List(ctx context.Context, identity int) (interface{}, error) {
	if identity == constant.StudentIdentity {
		var students []entity.Student
		err := mysql.DB.Model(&entity.Student{}).Find(&students).Error
		if err != nil {
			return nil, err
		}
		return students, nil
	} else if identity == constant.TeacherIdentity {
		var teachers []entity.Teacher
		err := mysql.DB.Model(&entity.Teacher{}).Find(&teachers).Error
		if err != nil {
			return nil, err
		}
		return teachers, nil
	}
	return nil, nil
}

func DeleteStudent(ctx context.Context, studentId uint) error {
	return mysql.DB.Delete(&entity.Student{}, studentId).Error
}

func DeleteTeacher(ctx context.Context, teacherId uint) error {
	return mysql.DB.Delete(&entity.Teacher{}, teacherId).Error
}

func ChangeAdminPassword(username string, password string) error {
	return mysql.DB.Model(&entity.Admin{}).Where("username = ?", username).Update("password", password).Error
}
