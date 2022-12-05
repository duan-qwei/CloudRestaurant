package controller

import (
	"CloudRestaurant/constant"
	"CloudRestaurant/model/reponse"
	"CloudRestaurant/model/request"
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
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
	err := c.ShouldBindJSON(req)
	if err != nil {
		reponse.ResponseMessageReturn(c, http.StatusOK, http.StatusOK, constant.BindArgsError)
		return
	}

	roleService.Update(c, req)
}
