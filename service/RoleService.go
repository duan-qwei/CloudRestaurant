package service

import (
	"CloudRestaurant/common"
	"CloudRestaurant/config"
	"CloudRestaurant/constant"
	"CloudRestaurant/model"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleService struct {
}

func (s *RoleService) Add(c *gin.Context, req request.RoleAddReq) {
	var role model.Role
	db := common.DB.Where("name", req.Name).Find(&role)
	if db != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	if db.RowsAffected > 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataExist)
		return
	}

	worker, _ := common.NewWorker(config.Conf.WorkId)
	role.Id = worker.GetId()
	role.Name = req.Name

	create := common.DB.Create(&role)
	if create != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (s *RoleService) Update(c *gin.Context, req request.RoleUpdateReq) {
	var role model.Role
	db := common.DB.First(&role, req.Id)
	if db.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	if db.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	update := common.DB.Model(&role).Select("name").Updates(&role)
	if update.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return

}
