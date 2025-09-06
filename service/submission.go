package service

import (
	"context"
	"errors"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetSubmissionByHomeworkIdAndStudentId(homeworkId uint, studentId uint) ([]entity.Submission, error) {
	var submissions []entity.Submission
	err := mysql.DB.Model(&entity.Submission{}).Where("homework_id = ? and student_id = ?", homeworkId, studentId).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func GetSubmissionByHomeworkId(homeworkId uint) ([]entity.Submission, error) {
	var submissions []entity.Submission
	err := mysql.DB.Model(&entity.Submission{}).Where("homework_id = ?", homeworkId).Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	return submissions, nil
}

func CreateSubmission(ctx context.Context, submission entity.Submission) error {
	err := mysql.DB.Model(&entity.Submission{}).Create(&submission).Error
	if err != nil {
		return err
	}
	return nil
}

func RemoveSubmission(ctx context.Context, submissionId uint, studentId uint) error {
	delSubmission := entity.Submission{}
	err := mysql.DB.Where("id = ?", submissionId).First(&delSubmission).Error
	if delSubmission.StudentId != studentId {
		return errors.New("submission student does not belong to homework")
	}
	err = mysql.DB.Model(&entity.Submission{}).Where("id = ?", submissionId).Delete(&entity.Submission{}).Error
	if err != nil {
		return err
	}
	return nil
}
