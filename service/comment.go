package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
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
	err := mysql.DB.Model(&entity.Comment{}).Create(&comment).Error
	return err
}

func RemoveComment(c context.Context, id uint) error {
	err := mysql.DB.Model(&entity.Comment{}).Delete(&entity.Comment{}, id).Error
	return err
}
