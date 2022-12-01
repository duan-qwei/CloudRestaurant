package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/constant"
	"CloudRestaurant/model"
	"CloudRestaurant/model/reponse"
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
		c.JSON(http.StatusOK, reponse.Response{
			Code:    http.StatusOK,
			Message: constant.ERROR,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reponse.Response{
		Code:    http.StatusOK,
		Message: constant.SUCCESS,
	})
	return
}

func (u *UserReq) SelectUserById(userId int64) (data interface{}) {
	var (
		result model.User
	)
	common.DB.First(&result, userId)
	return result
}
