package redis

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var RDB *redis.Client

func Init() {
	cfg, err := ini.Load("./init/project.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	host := cfg.Section("redis").Key("host").String()
	port := cfg.Section("redis").Key("port").String()
	password := cfg.Section("redis").Key("password").String()

	RDB = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port, // 哨兵或集群时这里写哨兵地址或集群任意节点
		Password: password,          // 如果没有密码留空
		DB:       0,                 // 默认 0 号库

		// 连接池配置（生产环境建议按需调整）
		PoolSize:     100, // 最大 socket 连接数
		MinIdleConns: 10,  // 最小空闲连接
		PoolTimeout:  30 * time.Second,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		panic("连接失败: " + err.Error())
	}
}
