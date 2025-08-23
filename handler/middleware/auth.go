package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"studyonline/constant"
	"studyonline/dao/redis"
)

func Auth(iden int) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		result, err := redis.RDB.Get(c, token).Result()
		if err != nil {
			log.Println("token认证失败", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			c.Abort()
		}
		resSplit := strings.Split(result, "_")
		userId := resSplit[0]
		identity := resSplit[1]

		if iden == constant.CommonIdentity {
			c.Set("userId", userId)
			c.Set("identity", identity)
			c.Next()
		}

		identityInt, err := strconv.Atoi(identity)
		if iden == identityInt {
			c.Set("userId", userId)
			c.Set("identity", identity)
			c.Next()
		} else {
			log.Println("token认证失败", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "请求失败",
			})
			c.Abort()
		}
	}
}
