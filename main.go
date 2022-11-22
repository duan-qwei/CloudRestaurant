package main

import (
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {

	config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	r := gin.Default()

	r.Run(config.AppHost + ":" + config.AppPort)
}
