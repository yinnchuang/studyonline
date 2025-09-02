package login

import (
	"net/http"
	"studyonline/constant"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func StudentLogin(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	success, token, err := service.Login(c, username, password, constant.StudentIdentity)
	if err != nil || !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"token":   token,
	})
}

func TeacherLogin(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	success, token, err := service.Login(c, username, password, constant.TeacherIdentity)
	if err != nil || !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"token":   token,
	})
	return
}

func AdminLogin(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	success, token, err := service.Login(c, username, password, constant.AdminIdentity)
	if err != nil || !success {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"token":   token,
	})
	return
}
