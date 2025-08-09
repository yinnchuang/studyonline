package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 登录
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	// 获取资源
	v1 := r.Group("/resource")
	{
		v1.GET("/file", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v1.GET("/dataset", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
	}
	r.Run(":8000")
}
