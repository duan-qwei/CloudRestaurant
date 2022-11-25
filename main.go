package main

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
	"CloudRestaurant/router"
	gin "github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//读取并解析配置文件
	config.InitConfigFile()

	//初始化mysql
	common.InitDataSource()

	//初始化redis
	common.Init()
	defer common.Close()

	router.Router(r)

	//监听端口
	server := config.Conf.Server
	r.Run(":" + server.HttpPort)
}
