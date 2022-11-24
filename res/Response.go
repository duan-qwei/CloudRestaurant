package res

import "github.com/gin-gonic/gin"

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		g.C.JSON(httpCode, Response{
			Code:    intCode,
			Message: "参数错误",
			Data:    data,
		})
	case string:
		strCode := errCode.(string)
		g.C.JSON(httpCode, Response{
			Code:    9999,
			Message: strCode,
			Data:    data,
		})
	}

	return
}
