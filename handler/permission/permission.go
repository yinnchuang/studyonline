package permission

import (
	"net/http"
	"strconv"
	"strings"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

type PermissionVO struct {
	UserId   uint `json:"user_id"`
	Identity int  `json:"identity"`
}

func ListPermissionsByDatasetId(c *gin.Context) {
	datasetIdStr := c.DefaultQuery("dataset_id", "0")
	datasetId, _ := strconv.Atoi(datasetIdStr)
	permissions, err := service.ListNeedAuthPermission(c, uint(datasetId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var permissionVOs []PermissionVO
	for _, permission := range permissions {
		spl := strings.Split(permission, "_")
		userId, _ := strconv.Atoi(spl[0])
		identity, _ := strconv.Atoi(spl[1])
		permissionVOs = append(permissionVOs, PermissionVO{
			UserId:   uint(userId),
			Identity: identity,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    permissionVOs,
	})
}

type RequestPermissionByDatasetIdDTO struct {
	DatasetId uint `json:"dataset_id"`
}

func RequestPermissionByDatasetId(c *gin.Context) {
	var requestPermissionByDatasetIdDTO RequestPermissionByDatasetIdDTO
	err := c.ShouldBindJSON(&requestPermissionByDatasetIdDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	err = service.SetNeedAuthPermission(c, requestPermissionByDatasetIdDTO.DatasetId, userId, identity)
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

type CreatePermissionDTO struct {
	UserId    uint `json:"user_id"`
	Identity  int  `json:"identity"`
	DatasetId uint `json:"dataset_id"`
}

func CreatePermission(c *gin.Context) {
	var createPermissionDTO CreatePermissionDTO
	if err := c.ShouldBindJSON(&createPermissionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	teacherId := c.GetUint("userId")
	err := service.CreatePermission(createPermissionDTO.UserId, createPermissionDTO.Identity,
		createPermissionDTO.DatasetId, teacherId)
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
