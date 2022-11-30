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
		c.JSON(http.StatusInternalServerError, reponse.Response{
			Code:    http.StatusInternalServerError,
			Message: "绑定错误",
			Data:    nil,
		})
		return
	}

	userService.Insert(c, &addUser)
}

// GetUerInfoById 根据id获取用户详情
func (userController *UserController) GetUerInfoById(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		c.JSON(http.StatusOK, reponse.Response{
			Code:    http.StatusBadRequest,
			Message: constant.UserIdIsBLANK,
		})
		return
	}
	id := com.StrTo(userId).MustInt64()
	data := userService.SelectUserById(id)

	c.JSON(http.StatusOK, reponse.Response{
		Code:    http.StatusOK,
		Message: constant.SUCCESS,
		Data:    data,
	})
	return
}
