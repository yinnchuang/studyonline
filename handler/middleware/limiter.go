package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 全局令牌桶
var globalLimiter *rate.Limiter
var once sync.Once

// 最大愿意让请求排队多久（0 表示必须立即拿到令牌）
const maxWait = 500 * time.Millisecond

// 初始化全局限流器：每秒 100 次，桶容量 200
func initGlobalLimiter() {
	once.Do(func() {
		globalLimiter = rate.NewLimiter(rate.Every(10*time.Millisecond), 200)
	})
}

// RateLimitGlobal 全局限速中间件，拿不到令牌时会延迟等待
func RateLimitGlobal() gin.HandlerFunc {
	initGlobalLimiter()

	return func(c *gin.Context) {
		// 尝试预约一个令牌
		r := globalLimiter.Reserve()
		delay := r.Delay()

		// 如果需要等待的时间超过阈值，直接取消并返回 429
		if delay > maxWait {
			r.Cancel() // 记得取消，把令牌还给桶
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "服务器繁忙，请稍后再试",
			})
			return
		}

		// 真正阻塞等待
		if delay > 0 {
			time.Sleep(delay)
		}

		c.Next()
	}
}
