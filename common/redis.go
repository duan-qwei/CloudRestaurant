package common

import (
	"CloudRestaurant/config"
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

// Init 初始化连接
func Init() {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Conf.RedisConfig.Host, config.Conf.RedisConfig.Port),
		Password: config.Conf.RedisConfig.Password,
		DB:       config.Conf.RedisConfig.DB,
		PoolSize: config.Conf.RedisConfig.PoolSize,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	log.Println("-------------redis初始化成功-------------")
}

func Close() {
	_ = client.Close()
}
