package user

import (
	"fmt"
	"net/http"
	"studyonline/constant"
	"studyonline/dao/redis"
	"studyonline/service"
	"studyonline/util"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	studentId := c.GetUint("userId")
	identity := c.GetInt("identity")
	if identity == constant.StudentIdentity || identity == constant.TeacherIdentity {
		userInfo, err := service.GetUserInfo(studentId, identity)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "请求成功",
				"data":    userInfo,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "请求失败",
	})
	return
}

type ChangePasswordDTO struct {
	Password string `json:"password"`
}

func ChangePassword(c *gin.Context) {
	var changePasswordDTO ChangePasswordDTO
	if err := c.ShouldBindJSON(&changePasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "请求失败",
		})
		return
	}

	userId := c.GetUint("userId")
	identity := c.GetInt("identity")
	bcryptPassword, _ := util.GetPwd(changePasswordDTO.Password)

	if !util.IsValidPassword(changePasswordDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码过于简单",
		})
		return
	}

	if identity == constant.StudentIdentity {
		err := service.ChangeStudentPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		userInfo, err := service.GetUserInfo(userId, identity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		cacheKey := fmt.Sprintf("change_password_%v_%v", userInfo.Username, constant.StudentIdentity)
		redis.RDB.Set(c, cacheKey, 1, time.Hour*24*7)
	} else if identity == constant.TeacherIdentity {
		err := service.ChangeTeacherPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		userInfo, err := service.GetUserInfo(userId, identity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
		cacheKey := fmt.Sprintf("change_password_%v_%v", userInfo.Username, constant.TeacherIdentity)
		redis.RDB.Set(c, cacheKey, 1, time.Hour*24*7)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
