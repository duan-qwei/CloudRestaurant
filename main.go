package main

import (
	"CloudRestaurant/config"
	"CloudRestaurant/datasource"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//读取并解析配置文件
	config.InitConfigFile()

	//初始化mysql
	datasource.InitDataSource()

	server := config.Conf.Server
	r.Run(server.HttpPort)
}
