package database

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/xjh22222228/open-erp/server/config"
)

var RedisClient *redis.Client

func RedisStart() {
	c := config.GlobalConfig.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.DB,
	})
	fmt.Printf("Redis 连接初始化成功: %s:%d\n", c.Host, c.Port)
}
