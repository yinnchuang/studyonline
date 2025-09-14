package score

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetScoreByStudentId(c *gin.Context) {
	studentId := c.GetUint("userId")
	results, err := service.GetScoreByStudentId(c, studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    results,
	})
}

func GetAllScore(c *gin.Context) {
	results, err := service.GetAllScore(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    results,
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
	score := entity.Score{
		StudentId:  createScoreDTO.StudentId,
		UsualScore: createScoreDTO.UsualScore,
		ExamScore:  createScoreDTO.ExamScore,
		FinalScore: createScoreDTO.FinalScore,
	}
	err := service.CreateScore(c, score)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
