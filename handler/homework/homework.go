package homework

import (
	"fmt"
	"net/http"
	"path/filepath"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

func ListHomework(c *gin.Context) {
	homeworks, err := service.GetAllHomework(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    homeworks,
	})
}

func UploadHomework(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	// 判断最大文件大小
	if file.Size > constant.MaxHomeworkSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件过大，不超过200M",
		})
		return
	}

	// 生成新文件名：时间戳 + 扩展名
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	newFileName = filepath.Join("./static/homework", newFileName)
	if err := c.SaveUploadedFile(file, newFileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传失败，稍后重试",
		})
		return
	}
	fileAbsPath, _ := filepath.Abs(newFileName)
	c.JSON(http.StatusOK, gin.H{
		"message":  "请求成功",
		"homework": fileAbsPath,
	})
}

type CreateHomeworkDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FilePath    string    `json:"file_path"`
	ExpireTime  time.Time `json:"expire_time"`
}

func CreateHomework(c *gin.Context) {
	createHomeworkDTO := CreateHomeworkDTO{}
	err := c.ShouldBindBodyWithJSON(&createHomeworkDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	homework := entity.Homework{
		Title:       createHomeworkDTO.Title,
		Description: createHomeworkDTO.Description,
		FilePath:    createHomeworkDTO.FilePath,
		ExpireTime:  createHomeworkDTO.ExpireTime,
	}
	err = service.CreateHomework(c, homework)
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

type RemoveHomeworkDTO struct {
	HomeworkId uint `json:"homework_id"`
}

func RemoveHomework(c *gin.Context) {
	removeHomeworkDTO := RemoveHomeworkDTO{}
	err := c.ShouldBindBodyWithJSON(&removeHomeworkDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveHomework(c, removeHomeworkDTO.HomeworkId)
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
