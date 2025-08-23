package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"studyonline/dao/redis"
)

func Auth(c *gin.Context) {
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
	c.Set("userId", userId)
	c.Set("identity", identity)
	c.Next()
}
