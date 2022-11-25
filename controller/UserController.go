package controller

import (
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"log"
)

var userService = service.User{}

type UserController struct {
}

func (e *UserController) InsertUser(c *gin.Context) {
	var addUser request.UserAddReq
	err := c.ShouldBindJSON(&addUser)

	if err != nil {
		log.Println("绑定失败")
	}

	userService.Insert(c, &addUser)
}
