package main

import (
	"CloudRestaurant/config"
	"CloudRestaurant/config/datasource"
	"CloudRestaurant/config/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//读取并解析配置文件
	config.InitConfigFile()

	//初始化mysql
	datasource.InitDataSource()

	//初始化redis
	redis.Init()
	defer redis.Close()

	//监听端口
	server := config.Conf.Server
	r.Run(":" + server.HttpPort)
}
