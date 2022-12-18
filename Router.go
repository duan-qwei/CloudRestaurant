package main

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
		roleController       = controller.RoleController{}
	)

	//用户管理
	adminRouter := engine.Group("/user/manage")
	{
		adminRouter.POST("/add", userManageController.InsertUser)
		adminRouter.DELETE("/delete/:id", userManageController.DeleteUserById)
		adminRouter.POST("/update", userManageController.Update)
		adminRouter.GET("/getInfo", userManageController.GetUerInfoById)
		adminRouter.GET("/getAllUser", userManageController.GetAllUser)
	}

	//用户前台
	userRouter := engine.Group("/user/interface")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
		userRouter.POST("/updateInfo", userController.UpdateInfoByUser)
		userRouter.GET("/getProfile/:id", userController.GetProfile)
	}

	//角色
	roleRouter := engine.Group("/role")
	{
		roleRouter.POST("/add", roleController.Add)
		roleRouter.POST("/update", roleController.Update)
		roleRouter.DELETE("/delete", roleController.Delete)
		roleRouter.GET("/getById/:id", roleController.GetById)
	}

	//积分
	pointRouter := engine.Group("/point")
	{
		pointRouter.GET("/add")
	}
}
