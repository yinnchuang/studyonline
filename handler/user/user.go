package user

import (
	"net/http"
	"studyonline/constant"
	"studyonline/service"

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
