package param

import (
	"CloudRestaurant/constant"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	error := c.Bind(form)

	if error != nil {
		log.Println(error.Error())
		return http.StatusInternalServerError, constant.InvalidParams
	}

	valid := validation.Validation{}

	ok, error := valid.Valid(form)
	if error != nil {
		return http.StatusInternalServerError, http.StatusBadRequest
	}

	if !ok {
		return http.StatusInternalServerError, constant.InvalidParams
	}

	return http.StatusOK, http.StatusBadRequest
}
