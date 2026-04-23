package service

import (
	"context"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDiscuss(c *gin.Context, id uint) (*entity.Discuss, error) {
	var discuss entity.Discuss
	err := mysql.DB.Model(&entity.Discuss{}).Where("id = ?", id).Find(&discuss).Error
	if err != nil {
		return nil, err
	}
	return &discuss, nil
}

func GetAllDiscusses(c context.Context) ([]entity.Discuss, error) {
	var discusses []entity.Discuss
	err := mysql.DB.Model(&entity.Discuss{}).Find(&discusses).Error
	if err != nil {
		return nil, err
	}
	return discusses, nil
}

func CreateDiscuss(c context.Context, discuss entity.Discuss) error {
	// 开启事务
	tx := mysql.DB.Begin()
	
	// 创建讨论
	if err := tx.Model(&entity.Discuss{}).Create(&discuss).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 如果是学生用户，增加评论数
	if discuss.Identity == constant.StudentIdentity {
		if err := tx.Model(&entity.Student{}).Where("id = ?", discuss.UserId).Update("comment_count", gorm.Expr("COALESCE(comment_count, 0) + ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	
	return tx.Commit().Error
}

func RemoveDiscuss(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.Discuss{}).Delete(&entity.Discuss{}, id).Error
	return err
}
