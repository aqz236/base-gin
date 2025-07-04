package cache

import (
	"base-gin/configs"
	"fmt"
	"log"
)

type RedisClient struct {
	config *configs.CacheConfig
}

func NewRedisClient(config *configs.Config) *RedisClient {
	client := &RedisClient{
		config: &config.Cache,
	}

	// 在真实项目中，这里会初始化 Redis 连接
	log.Printf("Redis配置: Host=%s, Port=%d, DB=%d",
		client.config.Host, client.config.Port, client.config.DB)

	return client
}

func (r *RedisClient) GetConnectionString() string {
	return fmt.Sprintf("%s:%d", r.config.Host, r.config.Port)
}

func (r *RedisClient) Close() error {
	// 在真实项目中，这里会关闭 Redis 连接
	log.Println("Redis连接已关闭")
	return nil
}
