package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 客户端桶的缓存
var visitorCache = map[string]*rate.Limiter{}

// 全局锁
var mu sync.RWMutex

func getLimiter(ip string) *rate.Limiter {
	mu.RLock()
	l, ok := visitorCache[ip]
	mu.RUnlock()

	if ok {
		return l
	}

	mu.Lock()
	defer mu.Unlock()
	l = rate.NewLimiter(rate.Every(25*time.Millisecond), 50)
	visitorCache[ip] = l
	return l
}

func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := getLimiter(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "请求过快",
			})
			return
		}
		c.Next()
	}
}
