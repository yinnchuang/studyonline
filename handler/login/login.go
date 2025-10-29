package login

import (
	"errors"
	"fmt"
	"net/http"
	"studyonline/constant"
	"studyonline/dao/redis"
	"studyonline/service"

	redis_ "github.com/redis/go-redis/v9"

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
	cacheKey := fmt.Sprintf("change_password_%v_%v", username, constant.StudentIdentity)
	changePassword, err := redis.RDB.Get(c, cacheKey).Int()
	if errors.Is(err, redis_.Nil) || changePassword == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message":         "请求成功",
			"token":           token,
			"change_password": 1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":         "请求成功",
		"token":           token,
		"change_password": -1,
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
	cacheKey := fmt.Sprintf("change_password_%v_%v", username, constant.TeacherIdentity)
	changePassword, err := redis.RDB.Get(c, cacheKey).Int()
	if errors.Is(err, redis_.Nil) || changePassword == 1 {
		c.JSON(http.StatusOK, gin.H{
			"message":         "请求成功",
			"token":           token,
			"change_password": 1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":         "请求成功",
		"token":           token,
		"change_password": -1,
	})
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
