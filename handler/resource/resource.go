package resource

import (
	"fmt"
	"net/http"
	"path/filepath"
	"studyonline/constant"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadResource(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	cover, err := c.FormFile("cover")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	// 判断最大文件大小
	if file.Size > constant.MaxResourceSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "文件过大",
		})
		return
	}
	if cover.Size > constant.MaxCoverSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "封面过大",
		})
		return
	}

	// 生成新文件名：时间戳 + 扩展名
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	newFileName = filepath.Join("./static/resource", newFileName)
	if err := c.SaveUploadedFile(file, newFileName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传失败，稍后重试",
		})
		return
	}
	ext = filepath.Ext(cover.Filename)
	newCoverName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	newCoverName = filepath.Join("./static/cover", newCoverName)
	if err := c.SaveUploadedFile(cover, newCoverName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传失败，稍后重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"file":    newFileName,
		"cover":   newCoverName,
	})
}

type CreateResourceDTO struct {
	Name        string `json:"name"`
	CategoryID  int    `json:"category_id"`
	Description string `json:"description"`
	FilePath    string `json:"file_path"`
	CoverPath   string `json:"cover_path"`
	UnitIds     []uint `json:"unit_ids"`
}

func CreateResource(c *gin.Context) {
	createResourceDTO := CreateResourceDTO{}
	err := c.ShouldBind(&createResourceDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.CreateResource(c, createResourceDTO.Name, createResourceDTO.CategoryID,
		createResourceDTO.Description, createResourceDTO.FilePath,
		createResourceDTO.CoverPath, createResourceDTO.UnitIds)
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
