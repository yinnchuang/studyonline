package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 哨兵或集群时这里写哨兵地址或集群任意节点
		Password: "",               // 如果没有密码留空
		DB:       0,                // 默认 0 号库

		// 连接池配置（生产环境建议按需调整）
		PoolSize:     100, // 最大 socket 连接数
		MinIdleConns: 10,  // 最小空闲连接
		PoolTimeout:  30 * time.Second,
	})
}
