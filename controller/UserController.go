package controller

import (
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

var userService = service.UserReq{}

type UserController struct {
}

func (userController *UserController) InsertUser(c *gin.Context) {
	var addUser request.UserAddReq
	err := c.ShouldBindJSON(&addUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "绑定失败")
		return
	}

	userService.Insert(c, &addUser)
}

func (userController *UserController) GetUerInfoById(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	data := userService.SelectUserById(c, id)
	log.Println(data)
}
