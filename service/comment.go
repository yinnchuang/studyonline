package service

import (
	"context"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"

	"gorm.io/gorm"
)

func GetCommentById(c context.Context, id uint) (*entity.Comment, error) {
	var comment entity.Comment
	err := mysql.DB.Where("id = ?", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func GetCommentByDiscussId(c context.Context, discussId uint) ([]entity.Comment, error) {
	var comments []entity.Comment
	err := mysql.DB.Model(&entity.Comment{}).Where("discuss_id = ?", discussId).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func CreateComment(c context.Context, comment entity.Comment) error {
	// 开启事务
	tx := mysql.DB.Begin()
	
	// 创建评论
	if err := tx.Model(&entity.Comment{}).Create(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 如果是学生用户，增加评论数
	if comment.Identity == constant.StudentIdentity {
		if err := tx.Model(&entity.Student{}).Where("id = ?", comment.UserId).Update("comment_count", gorm.Expr("COALESCE(comment_count, 0) + ?", 1)).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	
	// 获取被评论的 Discuss，增加被评论者的被评论数
	var discuss entity.Discuss
	if err := tx.Where("id = ?", comment.DiscussId).First(&discuss).Error; err == nil {
		// 如果 Discuss 作者是学生，增加被评论数
		if discuss.Identity == constant.StudentIdentity {
			if err := tx.Model(&entity.Student{}).Where("id = ?", discuss.UserId).Update("be_commented_count", gorm.Expr("COALESCE(be_commented_count, 0) + ?", 1)).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	
	return tx.Commit().Error
}

func RemoveComment(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.Comment{}).Delete(&entity.Comment{}, id).Error
	return err
}
