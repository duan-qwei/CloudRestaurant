package controller

import (
	"CloudRestaurant/constant"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

var userService = service.UserReq{}

type UserController struct {
}

// InsertUser 添加一个用户
func (userController *UserController) InsertUser(c *gin.Context) {
	var addUser request.UserAddReq
	err := c.ShouldBindJSON(&addUser)

	if err != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.BindDataError, err.Error())
		return
	}

	userService.Insert(c, &addUser)
}

// GetUerInfoById 根据id获取用户详情
func (userController *UserController) GetUerInfoById(c *gin.Context) {
	var req request.UserQueryInfoReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindDataError, err.Error())
		return
	}

	data := userService.SelectUserById(req.Id)
	reponse.ResponseReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS, data)
	return
}

// DeleteUserById 根据用户id删除
func (userController *UserController) DeleteUserById(c *gin.Context) {
	paramId := c.Param("id")
	if paramId == "" {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.ArgsIsBlank)
		return
	}

	userId := com.StrTo(paramId).MustInt64()
	userService.DeleteById(c, userId)
}

// Update 更新用户
func (userController *UserController) Update(c *gin.Context) {
	var req request.UserUpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusBadRequest, constant.BindArgsError, err.Error())
		return
	}
	userService.Update(c, req)
}
