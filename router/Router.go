package router

import (
	"CloudRestaurant/config"
	"CloudRestaurant/controller"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	gin.SetMode(config.Conf.Server.RunMode)

	var (
		userManageController = controller.UserManagerController{}
		userController       = controller.UserController{}
	)

	adminRouter := engine.Group("/user/manage")
	{
		adminRouter.POST("/add", userManageController.InsertUser)
		adminRouter.DELETE("/delete/:id", userManageController.DeleteUserById)
		adminRouter.POST("/update", userManageController.Update)
		adminRouter.GET("/getInfo", userManageController.GetUerInfoById)
	}

	userRouter := engine.Group("/user/interface")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login")
	}
}
