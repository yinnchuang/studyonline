package score

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studyonline/service"
)

func GetScoreByStudentId(c *gin.Context) {
	studentIdStr := c.DefaultQuery("student_id", "")
	if studentIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	studentId, _ := strconv.Atoi(studentIdStr)
	scores, err := service.GetScoreByStudentId(c, uint(studentId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"scores":  scores,
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
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"scores":  scores,
	})
}

func GetMeanScore(c *gin.Context) {
	scores, err := service.GetAllScore(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	meanScores := map[uint]float64{}
	scoreCounts := map[uint]uint{}
	for _, score := range scores {
		meanScores[score.StudentId] += float64(score.Score)
		scoreCounts[score.StudentId]++
	}
	for studentId, score := range meanScores {
		meanScores[studentId] = score / float64(scoreCounts[studentId])
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "请求成功",
		"mean_score": meanScores,
	})
}

type CreateScorePSO struct {
	StudentId uint `json:"student_id"`
	UnitId    uint `json:"unit_id"`
	Score     int  `json:"score"`
}

func CreateScore(c *gin.Context) {
	createScorePSO := CreateScorePSO{}
	if err := c.ShouldBindJSON(&createScorePSO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	teacherId := c.GetUint("teacher_id")
	err := service.CreateScore(c, teacherId, createScorePSO.StudentId, createScorePSO.UnitId, createScorePSO.Score)
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
