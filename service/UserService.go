package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
	"CloudRestaurant/constant"
	"CloudRestaurant/model"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"CloudRestaurant/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var (
	appConfig    config.AppConfig
	roleService  RoleService
	pointService PointService
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

func (u *UserReq) Register(c *gin.Context, register *request.UserRegisterAndLogin) {
	db, user := getOneByUsername(c, register.Username)
	if db.RowsAffected > 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataExist)
		return
	}

	user.Password = tools.GeneratePassword(register.Password)
	user.Username = register.Username
	user.Id = tools.GenerateNextId()
	save := common.DB.Create(&user)
	if err := save.Error; err != nil {
		log.Println(err.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.ERROR)
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) Login(c *gin.Context, req request.UserRegisterAndLogin) {
	db, user := getOneByUsername(c, req.Username)
	if db.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	_, err := tools.CheckPassword(user.Password, req.Password)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.PasswordIsNotRight)
		return
	}

	userLogin := model.UserLogin{
		Id:       user.Id,
		Phone:    user.Phone,
		Username: user.Username,
		Email:    user.Email,
		Picture:  user.Picture,
		RoleId:   user.RoleId,
	}

	if userLogin.RoleId != 0 {
		role := roleService.GetById(c, userLogin.RoleId)
		userLogin.RoleName = role.Name
	}

	point, err := pointService.getPointsByUserId(c, user.Id)
	if err != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.PasswordIsNotRight)
		return
	}
	if point != nil {
		userLogin.Points = point.Points
	}

	toStr, _ := json.Marshal(userLogin)
	common.RedisClient.Set(com.StrTo(userLogin.Id).String(), toStr, time.Second*100)
	reponse.ResponseReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS, &userLogin)
	return
}

func (u *UserReq) UpdateInfoByUser(c *gin.Context, req request.UserUpdateReq) {
	var user *model.User
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

	user.Email = req.Email
	user.Username = req.Username
	user.Phone = req.Phone
	user.Picture = req.Picture
	update := common.DB.Model(&user).Select("username", "phone", "email", "picture").Updates(&user)
	if err := update.Error; err != nil {
		log.Println(err.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.ERROR)
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (u *UserReq) GetAllUser(c *gin.Context) {

}

// GetProfile 获取用户的基本信息
func (u *UserReq) GetProfile(c *gin.Context, id int64) {
	user := getOneById(c, id)
	if user.RoleId != 0 {
		var role *model.Role
		db := common.DB.First(&role, user.RoleId)
		if db.Error != nil {
			reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.SqlError, db.Error)
			return
		}
		if role != nil {
			user.RoleName = role.Name
		}
	}
	reponse.ResponseReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS, user)
	return
}

// getOneByUsername 根据用户名获取用户
func getOneByUsername(c *gin.Context, username string) (db *gorm.DB, u *model.User) {
	var user *model.User
	db = common.DB.Where("username = ?", username).Find(&user)
	if error := db.Error; error != nil {
		log.Println(error.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.SqlError)
		return
	}
	return db, user
}

// getOneById 根据id获取单个用户
func getOneById(c *gin.Context, id int64) (u *model.User) {
	var user *model.User
	db := common.DB.First(&user, id)
	if error := db.Error; error != nil {
		log.Println(error.Error())
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusInternalServerError, constant.SqlError)
		return
	}

	return user
}
