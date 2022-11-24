package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid 绑定并校验参数
func BindAndValid(c *gin.Context, form interface{}) (int, string) {
	if err := c.Bind(form); err != nil {
		return http.StatusInternalServerError, "请求参数绑定有误！"
	}

	valid := validation.Validation{}

	result, err := valid.Valid(form)
	if err != nil || !result {
		return http.StatusInternalServerError, "请求参数验证有误！"
	}

	return http.StatusOK, "校验成功"
}
