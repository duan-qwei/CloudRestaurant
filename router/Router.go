package router

import (
	"CloudRestaurant/controller"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	userController := controller.UserController{}
	adminRouter := r.Group("/user")

	{
		adminRouter.POST("/add", userController.InsertUser)
	}
}
