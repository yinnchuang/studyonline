package permission

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"studyonline/service"
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
