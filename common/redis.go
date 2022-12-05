package common

import (
	"CloudRestaurant/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var (
	RedisClient *redis.Client
	Nil         = redis.Nil
)

// Init 初始化连接
func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Conf.RedisConfig.Host, config.Conf.RedisConfig.Port),
		Password: config.Conf.RedisConfig.Password,
		DB:       config.Conf.RedisConfig.DB,
		PoolSize: config.Conf.RedisConfig.PoolSize,
	})

	_, err := RedisClient.Ping().Result()

	if err != nil {
		panic(err)
	}

	log.Println("-------------redis初始化成功-------------")
}

func Close() {
	_ = RedisClient.Close()
}
