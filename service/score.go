package service

import (
	"context"
	"studyonline/dao/entity"
	"studyonline/dao/mysql"
)

func GetScoreByStudentId(c context.Context, studentId uint) (*entity.Score, error) {
	var score entity.Score
	err := mysql.DB.Model(&entity.Score{}).Where("student_id = ?", studentId).First(&score).Error
	if err != nil {
		return nil, err
	}
	return &score, nil
}

func GetAllScore(c context.Context) ([]entity.Score, error) {
	var scores []entity.Score
	err := mysql.DB.Model(&entity.Score{}).Find(&scores).Error
	if err != nil {
		return nil, err
	}
	return scores, nil
}

func CreateScore(c context.Context, score entity.Score) error {
	// 开启事务
	tx := mysql.DB.Begin()

	// 获取学生的活跃指标
	var student entity.Student
	var activityScore int = 0
	if err := tx.Where("id = ?", score.StudentId).First(&student).Error; err == nil {
		// 计算活跃度分数（三个活跃指标的总和*10，最高100）
		total := student.CommentCount + student.LikeCount + student.BeCommentedCount
		activityScore = total * 10
		if activityScore > 100 {
			activityScore = 100
		}
	}

	// 创建分数记录
	createData := entity.Score{
		StudentId:     score.StudentId,
		UsualScore:    score.UsualScore,
		ExamScore:     score.ExamScore,
		ActivityScore: activityScore,
	}

	if err := tx.Model(&entity.Score{}).Create(&createData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateScore(c context.Context, score entity.Score) error {
	// 开启事务
	tx := mysql.DB.Begin()

	// 获取学生的活跃指标
	var student entity.Student
	var activityScore int = 0
	if err := tx.Where("id = ?", score.StudentId).First(&student).Error; err == nil {
		// 计算活跃度分数（三个活跃指标的总和*10，最高100）
		total := student.CommentCount + student.LikeCount + student.BeCommentedCount
		activityScore = total * 10
		if activityScore > 100 {
			activityScore = 100
		}
	}

	// 更新分数记录，只更新必要的字段
	updateData := map[string]interface{}{
		"usual_score":    score.UsualScore,
		"exam_score":     score.ExamScore,
		"activity_score": activityScore,
	}

	if err := tx.Model(&entity.Score{}).Where("student_id = ?", score.StudentId).Updates(updateData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func ExistScore(c context.Context, studentID uint) (bool, error) {
	var count int64
	if err := mysql.DB.Model(&entity.Score{}).Where("student_id = ?", studentID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func DeleteScore(c context.Context, scoreId uint) error {
	return mysql.DB.Delete(&entity.Score{}, "id = ?", scoreId).Error
}
