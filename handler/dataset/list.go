package dataset

import (
	"fmt"
	"net/http"
	"strconv"
	"studyonline/constant"
	"studyonline/service"
	"studyonline/util"
	"time"

	"github.com/gin-gonic/gin"
)

type ListDatasetVO struct {
	ID                 uint      `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	Name               string    `json:"name"`
	CategoryID         int       `json:"category_id"`
	Description        string    `json:"description,omitempty"`
	Scale              string    `json:"scale"`
	Private            bool      `json:"private"`
	Url                string    `json:"url"`
	UploaderName       string    `json:"uploader_name"`
	UploaderUsername   string    `json:"uploader_username"`
	UploaderDepartment string    `json:"uploader_department"`
	DownloadTime       int       `json:"download_time"`
}

func ListDataset(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "-1")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit
	datasetWithLimitOffset, err := service.ListDatasetWithLimitOffset(c, limit, offset)

	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasetWithLimitOffset {
		uploader, err := service.GetTeacherInfo(item.TeacherId)
		if err != nil || uploader == nil {
			continue
		}

		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:                 item.ID,
			CreatedAt:          item.CreatedAt,
			Name:               item.Name,
			CategoryID:         item.CategoryID,
			Description:        item.Description,
			Scale:              item.Scale,
			Private:            item.Private,
			Url:                item.Url,
			UploaderName:       util.ProtectName(uploader.Name),
			UploaderUsername:   uploader.Username,
			UploaderDepartment: uploader.Department,
			DownloadTime:       item.DownloadTime,
		})
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountDataset(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
		"total":   total,
	})
}

func ListDatasetByCategory(c *gin.Context) {
	categoryStr := c.DefaultQuery("category", "-1")
	category, _ := strconv.Atoi(categoryStr)

	limitStr := c.DefaultQuery("limit", "3")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 非类别之一
	if !constant.IfDatasetCategory(category) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 展示特定种类
	datasetWithCategory, err := service.ListDatasetWithCategoryLimitOffset(c, limit, offset, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	total, err := service.CountDatasetWithCategory(c, category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	fmt.Println(datasetWithCategory)
	fmt.Println(total)
	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasetWithCategory {
		uploader, err := service.GetTeacherInfo(item.TeacherId)
		if err != nil || uploader == nil {
			continue
		}

		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:                 item.ID,
			CreatedAt:          item.CreatedAt,
			Name:               item.Name,
			CategoryID:         item.CategoryID,
			Description:        item.Description,
			Scale:              item.Scale,
			Private:            item.Private,
			Url:                item.Url,
			UploaderName:       util.ProtectName(uploader.Name),
			UploaderUsername:   uploader.Username,
			UploaderDepartment: uploader.Department,
			DownloadTime:       item.DownloadTime,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
		"total":   total,
	})
}

func ListDatasetByTeacherId(c *gin.Context) {
	userId := c.GetUint("userId")

	// 根据teacherId查询数据集
	datasets, err := service.ListDatasetByTeacherId(c, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasets {
		uploader, err := service.GetTeacherInfo(item.TeacherId)
		if err != nil || uploader == nil {
			continue
		}

		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:                 item.ID,
			CreatedAt:          item.CreatedAt,
			Name:               item.Name,
			CategoryID:         item.CategoryID,
			Description:        item.Description,
			Scale:              item.Scale,
			Private:            item.Private,
			Url:                item.Url,
			UploaderName:       util.ProtectName(uploader.Name),
			UploaderUsername:   uploader.Username,
			UploaderDepartment: uploader.Department,
			DownloadTime:       item.DownloadTime,
		})
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
	})
}

// SearchDatasetByKeyword 根据关键词搜索数据集（匹配名称或描述）
func SearchDatasetByKeyword(c *gin.Context) {
	// 1. 解析分页参数（与原有接口保持一致，默认limit=-1查全部，page=0起始）
	limitStr := c.DefaultQuery("limit", "-1")
	pageStr := c.DefaultQuery("page", "0")
	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)
	offset := page * limit

	// 2. 解析搜索关键词，空关键词直接返回参数错误
	keyword := c.Query("keyword")
	if keyword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "搜索关键词不能为空",
		})
		return
	}

	// 3. 调用service层查询匹配的数据集
	datasets, err := service.SearchDatasetByKeyword(c, limit, offset, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 4. 转换为VO结构（关联教师信息，与原有List接口逻辑一致）
	listDatasetVOs := []ListDatasetVO{}
	for _, item := range datasets {
		// 关联上传者（教师）信息，查询失败则跳过该条数据
		uploader, err := service.GetTeacherInfo(item.TeacherId)
		if err != nil || uploader == nil {
			continue
		}

		listDatasetVOs = append(listDatasetVOs, ListDatasetVO{
			ID:                 item.ID,
			CreatedAt:          item.CreatedAt,
			Name:               item.Name,
			CategoryID:         item.CategoryID,
			Description:        item.Description,
			Scale:              item.Scale,
			Private:            item.Private,
			Url:                item.Url,
			UploaderName:       util.ProtectName(uploader.Name),
			UploaderUsername:   uploader.Username,
			UploaderDepartment: uploader.Department,
			DownloadTime:       item.DownloadTime,
		})
	}

	// 5. 查询符合条件的总条数（用于分页计算）
	total, err := service.CountDatasetByKeyword(c, keyword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	// 6. 返回统一格式结果（与原有接口字段一致：message、data、total）
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listDatasetVOs,
		"total":   total,
	})
}
