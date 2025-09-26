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

type CreateResourceDTO struct {
	Name        string `form:"name"`
	CategoryID  int    `form:"category_id"`
	Description string `form:"description"`
	UnitIds     []uint `form:"unit_ids"`
}

func UploadAndCreateResource(c *gin.Context) {
	// 1. 获取并验证上传的文件
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

	// 2. 验证文件大小
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

	// 3. 保存文件到服务器（不暴露给前端的内部路径）
	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join("./static/resource", fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件保存失败",
		})
		return
	}

	ext = filepath.Ext(cover.Filename)
	coverName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	coverPath := filepath.Join("./static/cover", coverName)
	if err := c.SaveUploadedFile(cover, coverPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "封面保存失败",
		})
		return
	}

	// 4. 获取其他表单数据
	createResourceDTO := CreateResourceDTO{}
	if err := c.Bind(&createResourceDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	fmt.Println(createResourceDTO)
	// 5. 使用保存后的路径创建资源
	err = service.CreateResource(
		c,
		createResourceDTO.Name,
		createResourceDTO.CategoryID,
		createResourceDTO.Description,
		filePath,
		coverPath,
		createResourceDTO.UnitIds,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "请求失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}

func GetResourceCover(c *gin.Context) {
	resourceIdStr := c.DefaultQuery("resource_id", "0")
	resourceId, _ := strconv.Atoi(resourceIdStr)
	resource, err := service.GetResourceByID(c, uint(resourceId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if resource.CoverPath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 直接使用gin的File方法返回文件
	c.File(resource.CoverPath)
}

func GetResource(c *gin.Context) {
	resourceIdStr := c.DefaultQuery("resource_id", "0")
	resourceId, _ := strconv.Atoi(resourceIdStr)
	resource, err := service.GetResourceByID(c, uint(resourceId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if resource.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 直接使用gin的File方法返回文件
	c.File(resource.FilePath)
}
