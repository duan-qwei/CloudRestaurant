package controller

import (
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
)

var userService = service.User{}

type UserController struct {
}

func (e *UserController) InsertUser(c *gin.Context) {
	req := new(request.UserAddReq)
	userService.Insert(c, req)
}
