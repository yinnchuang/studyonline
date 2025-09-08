package score

import (
	"net/http"
	"strconv"
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

func GetScoreByHomeworkId(c *gin.Context) {
	homeworkIdStr := c.DefaultQuery("homework_id", "0")
	homeworkId, _ := strconv.Atoi(homeworkIdStr)
	results, err := service.GetScoreByHomeworkId(c, uint(homeworkId))
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
	HomeworkId uint `json:"homework_id"`
	Score      int  `json:"score"`
}

func CreateScore(c *gin.Context) {
	createScoreDTO := CreateScoreDTO{}
	if err := c.ShouldBindJSON(&createScoreDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	teacherId := c.GetUint("teacher_id")
	score := entity.Score{
		TeacherId:  teacherId,
		StudentId:  createScoreDTO.StudentId,
		HomeworkId: createScoreDTO.HomeworkId,
		Score:      createScoreDTO.Score,
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
