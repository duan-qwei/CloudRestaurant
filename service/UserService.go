package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/model"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserReq struct {
	U *request.UserAddReq
}

func (u *UserReq) Insert(c *gin.Context, req *request.UserAddReq) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	err := common.DB.Create(&user).Error
	if err != nil {
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

func (u *UserReq) SelectUserById(userId int64) (data interface{}) {
	var (
		result model.User
	)
	common.DB.First(&result, userId)
	return result
}
