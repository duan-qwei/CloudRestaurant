package router

import (
	"CloudRestaurant/controller"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	userController := controller.UserController{}
	adminRouter := engine.Group("/user")

	{
		adminRouter.POST("/add", userController.InsertUser)
	}
}
