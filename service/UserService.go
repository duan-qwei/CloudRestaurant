package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
	"CloudRestaurant/constant"
	"CloudRestaurant/model"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var appConfig config.AppConfig

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
	err := common.DB.Delete(&model.User{}, userId)
	if err.Error != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.ERROR, err.Error)
		return
	}
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) Update(c *gin.Context, req request.UserUpdateReq) {
	var user model.User
	selectOne := common.DB.First(&user, req.Id)
	if selectOne.Error != nil {
		log.Println(selectOne.Error.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.ERROR)
		return
	}

	if selectOne.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	if user.Password != req.Password {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.PasswordIsNotRight)
		return
	}
	user.Password = req.NewPassword
	user.Username = req.Username
	common.DB.Save(&user)
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) Register(c *gin.Context, register *request.UserRegister) {
	user := model.User{
		Username: register.Username,
		Phone:    register.Phone,
	}

	selectOne := common.DB.Where("username = ? OR phone = ?", user.Username, user.Phone).Find(&user)
	if error := selectOne.Error; error != nil {
		log.Println(error.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.SqlError)
		return
	}

	if selectOne.RowsAffected > 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataExist)
		return
	}

	//加密处理
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	worker, _ := common.NewWorker(int64(config.Conf.WorkId))
	user.Id = worker.GetId()
	save := common.DB.Save(&user)
	if err := save.Error; err != nil {
		log.Println(err.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.ERROR)
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}
