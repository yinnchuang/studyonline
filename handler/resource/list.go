package resource

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RandomListResource(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)
	c.String(200, fmt.Sprintf("hello %s\n", offset))
}
