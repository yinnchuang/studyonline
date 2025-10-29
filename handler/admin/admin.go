package admin

import (
	"net/http"
	"studyonline/service"
	"studyonline/util"

	"github.com/gin-gonic/gin"
)

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
	bcryptPassword, _ := util.GetPwd(changePasswordDTO.Password)

	if !util.IsValidPassword(changePasswordDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "密码过于简单",
		})
		return
	}

	err := service.ChangeAdminPassword("admin", string(bcryptPassword))
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
