package main

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
	gin "github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//读取并解析配置文件
	config.InitConfigFile()

	//初始化mysql
	common.InitDataSource()

	//初始化redis
	common.Init()
	defer common.Close()

	// 初始化Validator数据校验
	common.InitValidate()

	Router(engine)

	engine.Use()

	//监听端口
	server := config.Conf.Server
	engine.Run(":" + server.HttpPort)
}
