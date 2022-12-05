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
	var req request.UserRegisterAndLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindArgsError, err.Error())
		return
	}

	userService.Register(c, &req)
}

// Login 用户登陆
func (userController *UserController) Login(c *gin.Context) {
	var req request.UserRegisterAndLogin
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindArgsError, err.Error())
		return
	}

	userService.Login(c, req)
}

// UpdateInfoByUser 用户更新信息
func (userController *UserController) UpdateInfoByUser(c *gin.Context) {
	var req request.UserUpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindArgsError, err.Error())
		return
	}

	userService.UpdateInfoByUser(c, req)
}
