package announcement

import (
	"net/http"
	"studyonline/dao/entity"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func GetAllAnnouncements(c *gin.Context) {
	announcements, err := service.GetAllAnnouncement(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    announcements,
	})
}

type CreateAnnouncementPSO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateAnnouncement(c *gin.Context) {
	createAnnouncementPSO := CreateAnnouncementPSO{}
	err := c.ShouldBindBodyWithJSON(&createAnnouncementPSO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	announcement := entity.Announcement{
		Title:       createAnnouncementPSO.Title,
		Description: createAnnouncementPSO.Description,
	}
	err = service.CreateAnnouncement(c, announcement)
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

type RemoveAnnouncementPSO struct {
	Id uint `json:"id"`
}

func RemoveAnnouncement(c *gin.Context) {
	removeAnnouncementPSO := RemoveAnnouncementPSO{}
	err := c.ShouldBindBodyWithJSON(&removeAnnouncementPSO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	err = service.RemoveAnnouncement(c, removeAnnouncementPSO.Id)
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
