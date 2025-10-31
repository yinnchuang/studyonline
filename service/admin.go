package service

import (
	"context"
	"errors"
	"fmt"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
	"studyonline/dao/redis"
	"time"
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

func BatchImportStudents(ctx context.Context, students []entity.Student) error {
	if len(students) == 0 {
		return errors.New("没有可导入的学生数据")
	}

	// 开启事务
	tx := mysql.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %v", tx.Error)
	}

	for _, stu := range students {
		// 执行插入
		if err := tx.Create(&stu).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("导入失败: %v", err)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("事务提交失败: %v", err)
	}

	for _, stu := range students {
		cacheKey := fmt.Sprintf("change_password_%v_%v", stu.Username, constant.StudentIdentity)
		redis.RDB.Set(ctx, cacheKey, -1, time.Hour*24*60)
	}

	return nil
}

func BatchImportTeachers(ctx context.Context, teachers []entity.Teacher) error {
	if len(teachers) == 0 {
		return errors.New("没有可导入的学生数据")
	}

	// 开启事务
	tx := mysql.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("开启事务失败: %v", tx.Error)
	}

	for _, tea := range teachers {
		// 执行插入
		if err := tx.Create(&tea).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("导入失败: %v", err)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("事务提交失败: %v", err)
	}

	for _, tea := range teachers {
		cacheKey := fmt.Sprintf("change_password_%v_%v", tea.Username, constant.TeacherIdentity)
		redis.RDB.Set(ctx, cacheKey, -1, time.Hour*24*60)
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
