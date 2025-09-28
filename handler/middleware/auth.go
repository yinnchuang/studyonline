package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"studyonline/constant"
	"studyonline/dao/redis"
	"studyonline/service"

	"github.com/gin-gonic/gin"
)

func Auth(iden int) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		log.Println(token)
		result, err := redis.RDB.Get(c, token).Result()
		if err != nil {
			log.Println("token认证失败", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			c.Abort()
		}
		resSplit := strings.Split(result, "_")
		userIdStr := resSplit[0]
		identityStr := resSplit[1]
		userId_, _ := strconv.Atoi(userIdStr)
		userId := uint(userId_)
		identity, _ := strconv.Atoi(identityStr)

		info, err := service.GetUserInfo(userId, identity)
		if err != nil {
			log.Println("token认证失败", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
		}

		log.Println("id为", userId, "的用户，身份为", identity, "正在请求", c.FullPath())

		// 如果是学生或老师，设置
		if iden == constant.StudentIdentity || iden == constant.TeacherIdentity {
			c.Set("name", info.Name)
			c.Set("department", info.Department)
		}

		if iden == constant.CommonIdentity {
			c.Set("userId", userId)
			c.Set("identity", identity)
			c.Next()
			return
		}

		if iden == constant.StaffIdentity {
			if identity == constant.AdminIdentity || identity == constant.TeacherIdentity {
				c.Set("userId", userId)
				c.Set("identity", identity)
				c.Next()
				return
			}
		}

		if iden == identity {
			c.Set("userId", userId)
			c.Set("identity", identity)
			c.Next()
			return
		} else {
			log.Println("token认证失败", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			c.Abort()
		}
	}
}
