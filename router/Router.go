package router

import (
	"CloudRestaurant/config"
	"CloudRestaurant/controller"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	gin.SetMode(config.Conf.Server.RunMode)
	userController := controller.UserController{}
	adminRouter := engine.Group("/user")

	{
		adminRouter.POST("/add", userController.InsertUser)
	}
}
