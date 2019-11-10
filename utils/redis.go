package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"sync"
)

type DBRedis struct {
	RedisClient *redis.Client
}

var (
	onceDbRedis     sync.Once
	instanceDBRedis *DBRedis
)

// This connection for redis
func GetInstanceRedis() *redis.Client {
	onceDbRedis.Do(func() {
		redisInfo := Config.Database.Redis
		logs := fmt.Sprintf("[INFO] Connected to REDIS Host = %s | Db = %d", redisInfo.Host, redisInfo.DB)

		clientConnection := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", redisInfo.Host, redisInfo.Port),
			Password: redisInfo.Password,
			DB:       redisInfo.DB,
		})

		if _, err := clientConnection.Ping().Result(); err != nil {
			logs = "[ERROR] Failed to connect to Redis. Config=" + redisInfo.Host
			log.Fatalln(logs)
		}
		fmt.Println(logs)
		instanceDBRedis = &DBRedis{RedisClient: clientConnection}
	})
	return instanceDBRedis.RedisClient
}
