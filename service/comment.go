package service

import (
	"context"
	"errors"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

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

func RemoveComment(c context.Context, id uint, userId uint) error {
	var comment entity.Comment
	err := mysql.DB.Model(&entity.Comment{}).Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return err
	}
	if comment.UserId != userId {
		return errors.New("dont match user id")
	}
	err = mysql.DB.Model(&entity.Comment{}).Delete(&entity.Comment{}, id).Error
	return err
}
