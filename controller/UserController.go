package controller

import (
	"CloudRestaurant/constant"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
}

// Register 注册用户
func (userController *UserController) Register(c *gin.Context) {
	var req request.UserRegister
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindArgsError, err.Error())
		return
	}

	userService.Register(c, &req)
}

func login(userController *UserController) (c *gin.Context) {

	return nil
}
