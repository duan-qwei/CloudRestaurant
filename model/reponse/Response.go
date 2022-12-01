package reponse

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

func ResponseErrorReturn(g *gin.Context, httCode, errorCode int, message string, error interface{}) {
	g.JSON(httCode, Response{
		Code:    errorCode,
		Message: message,
		Error:   error,
	})
}

func ResponseMessageReturn(g *gin.Context, httCode, errorCode int, message string) {
	g.JSON(httCode, Response{
		Code:    errorCode,
		Message: message,
	})
}

func ResponseReturn(g *gin.Context, httCode, errorCode int, message string, data interface{}) {
	g.JSON(httCode, Response{
		Code:    errorCode,
		Message: message,
		Data:    data,
	})
}
func (c *Gin) ResponsePage(httpCode int, errCode interface{}, data interface{}, total, totalPage int) {
	intCode := errCode.(int)
	c.C.JSON(httpCode, ResponsePage{
		Code:      intCode,
		Message:   "",
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
}
