package permission

import (
	"net/http"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

// 查看自己的申请记录
type ListRequestByUserIdVO struct {
	ID                  uint   `json:"id"`
	RequesterName       string `json:"requester_name"`
	RequesterDepartment string `json:"requester_department"`
	DatasetId           uint   `json:"dataset_id"`
	DatasetName         string `json:"dataset_name"`
	Reason              string `json:"reason"`
	Status              int    `json:"status"`
}

func ListRequestByUserId(c *gin.Context) {
	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	requests, err := service.ListRequestByUserId(c, userId, identity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var listRequestByUserIdVOs []ListRequestByUserIdVO
	for _, item := range requests {
		user, err1 := service.GetUserInfo(item.UserID, item.Identity)
		dataset, err2 := service.GetDatasetByID(c, item.DatasetId)
		if err1 != nil || err2 != nil {
			continue
		}
		listRequestByUserIdVO := ListRequestByUserIdVO{
			ID:                  item.ID,
			RequesterName:       user.Name,
			RequesterDepartment: user.Department,
			DatasetId:           item.DatasetId,
			DatasetName:         dataset.Name,
			Reason:              item.Reason,
			Status:              item.Status,
		}

		listRequestByUserIdVOs = append(listRequestByUserIdVOs, listRequestByUserIdVO)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listRequestByUserIdVOs,
	})
}

// 审批记录
type ListRequestByTeacherIdVO struct {
	ID                  uint   `json:"id"`
	RequesterName       string `json:"requester_name"`
	RequesterDepartment string `json:"requester_department"`
	DatasetId           uint   `json:"dataset_id"`
	DatasetName         string `json:"dataset_name"`
	Reason              string `json:"reason"`
	Status              int    `json:"status"`
}

func ListRequestByTeacherId(c *gin.Context) {
	teacherId := c.GetUint("userId")
	requests, err := service.ListRequestByTeacherId(c, teacherId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	var listRequestByTeacherIdVOs []ListRequestByTeacherIdVO
	for _, item := range requests {
		user, err1 := service.GetUserInfo(item.UserID, item.Identity)
		dataset, err2 := service.GetDatasetByID(c, item.DatasetId)
		if err1 != nil || err2 != nil {
			continue
		}
		listRequestByTeacherIdVO := ListRequestByTeacherIdVO{
			ID:                  item.ID,
			RequesterName:       user.Name,
			RequesterDepartment: user.Department,
			DatasetId:           item.DatasetId,
			DatasetName:         dataset.Name,
			Reason:              item.Reason,
			Status:              item.Status,
		}

		listRequestByTeacherIdVOs = append(listRequestByTeacherIdVOs, listRequestByTeacherIdVO)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    listRequestByTeacherIdVOs,
	})
}

// 发起申请
type RequestPermissionByDatasetIdDTO struct {
	DatasetId uint   `json:"dataset_id"`
	Reason    string `json:"reason"`
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
	dataset, err := service.GetDatasetByID(c, requestPermissionByDatasetIdDTO.DatasetId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.CreateRequest(c, requestPermissionByDatasetIdDTO.DatasetId, userId, identity,
		requestPermissionByDatasetIdDTO.Reason, dataset.TeacherId)
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

// 同意申请
type AgreePermissionDTO struct {
	RequestId uint `json:"request_id"`
}

func AgreePermission(c *gin.Context) {
	var agreePermissionDTO AgreePermissionDTO
	if err := c.ShouldBindJSON(&agreePermissionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	request, err := service.GetRequestById(c, agreePermissionDTO.RequestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	teacherId := c.GetUint("userId")

	err = service.AgreeRequest(c, agreePermissionDTO.RequestId, request.UserID, request.Identity,
		request.DatasetId, teacherId)

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

// 拒绝申请
type DisagreePermissionDTO struct {
	RequestId uint `json:"request_id"`
}

func DisagreePermission(c *gin.Context) {
	var disagreePermissionDTO DisagreePermissionDTO
	if err := c.ShouldBindJSON(&disagreePermissionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	request, err := service.GetRequestById(c, disagreePermissionDTO.RequestId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	teacherId := c.GetUint("userId")

	err = service.DisagreeRequest(c, disagreePermissionDTO.RequestId, request.UserID, request.Identity,
		request.DatasetId, teacherId)

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
