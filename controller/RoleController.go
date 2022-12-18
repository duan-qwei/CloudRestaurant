package controller

import (
	"CloudRestaurant/constant"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

var roleService = service.RoleService{}

type RoleController struct {
}

func (role *RoleController) Add(c *gin.Context) {
	var req request.RoleAddReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.BindArgsError)
		return
	}

	roleService.Add(c, req)
}

func (role *RoleController) Update(c *gin.Context) {
	var req request.RoleUpdateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err.Error())
		reponse.ResponseErrorReturn(c, http.StatusOK, http.StatusOK, constant.BindArgsError, err.Error())
		return
	}

	roleService.Update(c, req)
}

func (role *RoleController) Delete(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.ArgsIsBlank)
		return
	}

	roleService.Delete(c, com.StrTo(id).MustInt64())
}

func (role *RoleController) GetById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.ArgsIsBlank)
		return
	}

	selectById := roleService.GetById(c, com.StrTo(id).MustInt64())
	reponse.ResponseReturn(c, http.StatusOK, http.StatusOK, constant.SUCCESS, selectById)
	return
}
