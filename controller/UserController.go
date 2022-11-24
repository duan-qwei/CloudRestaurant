package controller

import (
	"CloudRestaurant/app"
	"CloudRestaurant/res"
	"CloudRestaurant/vo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
}

func (e *userController) insertUser(c *gin.Context) {
	var (
		model vo.SysUser
		appG  = res.Gin{C: c}
	)

	code, message := app.BindAndValid(c, &model)

	if code != http.StatusOK {
		appG.Response(code, message, nil)
		return
	}

	//userService := service.User{
	//	M: &model,
	//}

}
