package main

import (
	"CloudRestaurant/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//读取并解析配置文件
	config.InitConfigFile()
	r.Run()
}
