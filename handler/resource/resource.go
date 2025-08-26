package resource

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"studyonline/constant"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadResource(c *gin.Context) {
	resource, err := c.FormFile("resource")
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
	if resource.Size > constant.MaxResourceSize {
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
	ext := filepath.Ext(resource.Filename)
	newResourceName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	newResourceName = filepath.Join("./static/resource", newResourceName)
	if err := c.SaveUploadedFile(resource, newResourceName); err != nil {
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
		"message":  "请求成功",
		"resource": newResourceName,
		"cover":    newCoverName,
	})
}

type CreateResourceRequest struct {
	Name         string `json:"name"`
	CategoryID   int    `json:"category_id"`
	Description  string `json:"description"`
	ResourcePath string `json:"resource_path"`
	CoverPath    string `json:"cover_path"`
}

func CreateResource(c *gin.Context) {
	createResourceRequest := CreateResourceRequest{}
	err := c.ShouldBind(&createResourceRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.CreateResource(c, createResourceRequest.Name,
		createResourceRequest.CategoryID, createResourceRequest.Description,
		createResourceRequest.ResourcePath, createResourceRequest.CoverPath)
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

func GetResourceFile(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)
	c.String(200, fmt.Sprintf("hello %s\n", offset))
}

func GetResourceDataset(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)
	c.String(200, fmt.Sprintf("hello %s\n", offset))
}
