package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/constant"
	"CloudRestaurant/model"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"log"
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

	checkName := common.DB.Where("name=?", user.Username).Find(&user)

	if checkName != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.UserNameRepeat)
		return
	}
	err := common.DB.Create(&user).Error
	if err != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.ERROR, err.Error())
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) SelectUserById(userId int64) (data interface{}) {
	var (
		result model.User
	)
	common.DB.First(&result, userId)
	return result
}

func (u *UserReq) DeleteById(c *gin.Context, userId int64) {
	user := model.User{Id: userId}
	err := common.DB.Delete(&user).Error
	if err != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.ERROR, err.Error())
		return
	}
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) Update(c *gin.Context, req request.UserUpdateReq) {
	user := model.User{
		Id: req.Id,
	}
	selectOne := common.DB.Find(&user)
	if selectOne.Error != nil {
		log.Println(selectOne.Error.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.ERROR)
		return
	}

	if selectOne.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	user.Password = req.NewPassword
	user.Username = req.Username
	common.DB.Save(&user)
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}
