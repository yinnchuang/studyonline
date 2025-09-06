package submission

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"
	"time"
)

func ListSubmissionByHomeworkId(c *gin.Context) {
	homeworkIdStr := c.DefaultQuery("homework_id", "0")
	homeworkId, _ := strconv.Atoi(homeworkIdStr)
	identity := c.GetInt("identity")
	if identity == constant.StudentIdentity { // 如果是学生
		studentId := c.GetUint("userId")
		submissions, err := service.GetSubmissionByHomeworkIdAndStudentId(uint(homeworkId), studentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    submissions,
		})
	} else { // 如果是老师
		submissions, err := service.GetSubmissionByHomeworkId(uint(homeworkId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"data":    submissions,
		})
	}

}

func UploadSubmission(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	// 判断最大文件大小
	if file.Size > constant.MaxSubmissionSize {
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
		"message":    "请求成功",
		"submission": fileAbsPath,
	})
}

type CreateSubmissionDTO struct {
	HomeworkId  uint   `json:"homework_id"`
	FilePath    string `json:"file_path"`
	Description string `json:"description"`
}

func CreateSubmission(c *gin.Context) {
	createSubmissionDTO := CreateSubmissionDTO{}
	err := c.ShouldBindBodyWithJSON(&createSubmissionDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	homework, err := service.GetHomeworkById(c, createSubmissionDTO.HomeworkId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if homework.ExpireTime.Unix() < time.Now().Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "提交超时",
		})
		return
	}
	studentId := c.GetUint("userId")
	submission := entity.Submission{
		StudentId:   studentId,
		HomeworkId:  createSubmissionDTO.HomeworkId,
		FilePath:    createSubmissionDTO.FilePath,
		Description: createSubmissionDTO.Description,
	}
	err = service.CreateSubmission(c, submission)
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

type RemoveSubmissionDTO struct {
	SubmissionId uint `json:"submission_id"`
}

func RemoveSubmission(c *gin.Context) {
	removeSubmissionDTO := RemoveSubmissionDTO{}
	err := c.ShouldBindBodyWithJSON(&removeSubmissionDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	studentId := c.GetUint("userId")
	err = service.RemoveSubmission(c, removeSubmissionDTO.SubmissionId, studentId)
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
