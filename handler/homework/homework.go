package homework

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"studyonline/constant"
	"studyonline/dao/entity"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

type ListHomeworkVO struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireTime  time.Time `json:"expire_time"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
}

func ListHomework(c *gin.Context) {
	homeworks, err := service.GetAllHomework(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var listHomeworkVOs []ListHomeworkVO
	for _, item := range homeworks {
		listHomeworkVOs = append(listHomeworkVOs, ListHomeworkVO{
			ID:          item.ID,
			CreatedAt:   item.CreatedAt,
			ExpireTime:  item.ExpireTime,
			Title:       item.Title,
			Description: item.Description,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listHomeworkVOs,
	})
}

type CreateHomeworkDTO struct {
	Title       string    `form:"title"`
	Description string    `form:"description"`
	ExpireTime  time.Time `form:"expire_time"`
}

func UploadAndCreateHomework(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if file != nil && file.Size > constant.MaxHomeworkSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件过大",
		})
		return
	}

	now := time.Now()
	yearMonth := now.Format("200601")

	filePath := ""

	if file != nil {
		ext := filepath.Ext(file.Filename)
		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filePath = filepath.Join(fmt.Sprintf("./static/homework/%s", yearMonth), fileName)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "文件保存失败",
			})
			return
		}
	}

	createHomeworkDTO := CreateHomeworkDTO{}
	err = c.ShouldBind(&createHomeworkDTO)
	fmt.Println(createHomeworkDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	homework := entity.Homework{
		Title:       createHomeworkDTO.Title,
		Description: createHomeworkDTO.Description,
		FilePath:    filePath,
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

func GetHomework(c *gin.Context) {
	homeworkIdStr := c.DefaultQuery("homeworkId", "0")
	homeworkId, _ := strconv.Atoi(homeworkIdStr)

	homework, err := service.GetHomeworkById(c, uint(homeworkId))

	if err != nil || homework.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	c.File(homework.FilePath)
	return
}
