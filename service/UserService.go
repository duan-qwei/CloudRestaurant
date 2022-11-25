package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/model"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	U *request.UserAddReq
}

func (u *User) Insert(c *gin.Context, req *request.UserAddReq) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := common.DB.Create(&user).Error
	if err != nil {
		log.Println("创建用户失败", err)
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusInternalServerError,
			"msg":   "创建用户失败",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "创建用户成功",
	})
}
