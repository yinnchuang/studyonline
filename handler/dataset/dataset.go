package dataset

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"studyonline/constant"
	"studyonline/service"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateDatasetDTO struct {
	Name        string `form:"name"`
	CategoryID  int    `form:"category_id"`
	Description string `form:"description"`
	Scale       string `form:"scale"`
	Private     bool   `form:"private"`
	Url         string `form:"url"`
}

func UploadAndCreateDataset(c *gin.Context) {
	// 1. 获取并验证上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据集文件获取失败",
		})
		return
	}

	cover, err := c.FormFile("cover")

	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "封面图片获取失败",
		})
		return
	}

	// 2. 验证文件大小
	if file.Size > constant.MaxResourceSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "数据集文件过大",
		})
		return
	}

	if cover != nil && cover.Size > constant.MaxCoverSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "封面图片过大",
		})
		return
	}

	// 3. 保存文件到服务器（内部路径，不暴露给前端）
	now := time.Now()
	yearMonth := now.Format("200601")

	ext := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(fmt.Sprintf("./static/dataset/%s", yearMonth), fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "数据集文件保存失败",
		})
		return
	}

	coverPath := "./static/cover/cover.png"

	if cover != nil {
		ext = filepath.Ext(cover.Filename)
		coverName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		coverPath = filepath.Join(fmt.Sprintf("./static/cover/%s", yearMonth), coverName)
		if err := c.SaveUploadedFile(cover, coverPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "封面图片保存失败",
			})
			return
		}
	}

	// 4. 获取并验证其他表单数据
	datasetDTO := CreateDatasetDTO{}
	if err := c.ShouldBind(&datasetDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数解析失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	// 5. 调用服务层创建数据集记录
	dataset, err := service.CreateDataset(
		c,
		datasetDTO.Name,
		datasetDTO.CategoryID,
		datasetDTO.Description,
		filePath,  // 内部使用的文件路径
		coverPath, // 内部使用的封面路径
		datasetDTO.Scale,
		userId,
		datasetDTO.Private,
		datasetDTO.Url,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.CreatePermission(c, userId, identity, dataset.ID, dataset.TeacherId)
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

func GetDatasetCover(c *gin.Context) {
	datasetIdStr := c.DefaultQuery("dataset_id", "0")
	datasetId, _ := strconv.Atoi(datasetIdStr)
	dataset, err := service.GetDatasetByID(c, uint(datasetId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if dataset.CoverPath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 直接使用gin的File方法返回文件
	c.File(dataset.CoverPath)
}

func GetDataset(c *gin.Context) {
	datasetIdStr := c.DefaultQuery("dataset_id", "0")
	datasetId, _ := strconv.Atoi(datasetIdStr)
	dataset, err := service.GetDatasetByID(c, uint(datasetId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	if dataset.FilePath == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 非私有（公开）数据集
	if dataset.Private == false {
		service.PlusDatasetDownloadTime(c, uint(datasetId))

		// 日志
		name := c.GetString("name")
		department := c.GetString("department")
		service.AddLog(c, name, department, dataset.Name)

		c.File(dataset.FilePath)
		return
	}

	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if service.IfUserHasDatasetPermission(userId, identity, uint(datasetId)) {
		service.PlusDatasetDownloadTime(c, uint(datasetId))

		// 日志
		name := c.GetString("name")
		department := c.GetString("department")
		service.AddLog(c, name, department, dataset.Name)

		c.File(dataset.FilePath)
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "无下载权限，请申请权限",
		})
		return
	}

}
