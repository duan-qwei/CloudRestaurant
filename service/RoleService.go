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

var role = model.Role{}

type RoleService struct {
}

func (s *RoleService) Add(c *gin.Context, req request.RoleAddReq) {
	db := common.DB.Where("name", req.Name).Find(&role)
	if db.Error != nil {
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
	if create.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (s *RoleService) Update(c *gin.Context, req request.RoleUpdateReq) {
	db := common.DB.First(&role, req.Id)
	if db.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	if db.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	role.Name = req.Name
	update := common.DB.Model(&role).Select("name").Updates(&role)
	if update.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return

}

func (s *RoleService) Delete(c *gin.Context, id int64) {
	db := common.DB.First(&role, id)
	if db.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}

	delete := common.DB.Delete(&role)
	if delete.Error != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SqlError)
		return
	}
	reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS)
	return
}

func (s *RoleService) GetById(c *gin.Context, id int64) (role *model.Role) {
	db := common.DB.First(&role, id)
	if db.Error != nil {
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.SqlError, db.Error)
		return
	}

	if db.RowsAffected == 0 {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.DataIsNull)
		return
	}

	return role
}
