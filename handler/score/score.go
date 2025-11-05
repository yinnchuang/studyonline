package score

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

type ScoreVO struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	UserName   string `json:"user_name"`
	Department string `json:"department"`
	UsualScore int    `json:"usual_score"`
	ExamScore  int    `json:"exam_score"`
	FinalScore int    `json:"final_score"`
}

func GetScoreByStudentId(c *gin.Context) {
	studentId := c.GetUint("userId")
	score, err := service.GetScoreByStudentId(c, studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	studentInfo, err := service.GetStudentInfo(studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	scoreVO := ScoreVO{
		ID:         score.ID,
		Name:       studentInfo.Name,
		UserName:   studentInfo.Username,
		Department: studentInfo.Department,
		UsualScore: score.UsualScore,
		ExamScore:  score.ExamScore,
		FinalScore: score.FinalScore,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    scoreVO,
	})
}

func GetAllScore(c *gin.Context) {
	scores, err := service.GetAllScore(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var scoreVOs []ScoreVO
	for _, item := range scores {
		studentInfo, err := service.GetStudentInfo(item.StudentId)
		if err != nil || studentInfo == nil {
			continue
		}
		scoreVOs = append(scoreVOs, ScoreVO{
			ID:         item.ID,
			Name:       studentInfo.Name,
			UserName:   studentInfo.Username,
			Department: studentInfo.Department,
			UsualScore: item.UsualScore,
			ExamScore:  item.ExamScore,
			FinalScore: item.FinalScore,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    scoreVOs,
	})
}

type CreateScoreDTO struct {
	StudentId  uint `json:"student_id"`
	UsualScore int  `json:"usual_score" gorm:"not null"`
	ExamScore  int  `json:"exam_score" gorm:"not null"`
	FinalScore int  `json:"final_score" gorm:"not null"`
}

func CreateScore(c *gin.Context) {
	createScoreDTO := CreateScoreDTO{}
	if err := c.ShouldBindJSON(&createScoreDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	exist, err := service.ExistScore(c, createScoreDTO.StudentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	score := entity.Score{
		StudentId:  createScoreDTO.StudentId,
		UsualScore: createScoreDTO.UsualScore,
		ExamScore:  createScoreDTO.ExamScore,
		FinalScore: createScoreDTO.FinalScore,
	}

	if exist {
		err := service.UpdateScore(c, score)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	} else {
		err := service.CreateScore(c, score)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
