package user

import (
	"net/http"
	"studyonline/constant"
	"studyonline/service"
	"studyonline/util"

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

	if identity == constant.StudentIdentity {
		err := service.ChangeStudentPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	} else if identity == constant.TeacherIdentity {
		err := service.ChangeTeacherPassword(userId, string(bcryptPassword))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			return
		}
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
