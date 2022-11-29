package reponse

import "github.com/gin-gonic/gin"

type Gin struct {
	c *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponsePage struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
}

func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	switch errCode.(type) {
	case int:
		intCode := errCode.(int)
		g.c.JSON(httpCode, Response{
			Code:    intCode,
			Message: "",
			Data:    data,
		})
	case string:
		stringCode := errCode.(string)
		g.c.JSON(httpCode, Response{
			Code:    500,
			Message: stringCode,
			Data:    data,
		})
	}
	return
}

func (g *Gin) ResponsePage(httpCode int, errCode interface{}, data interface{}, total, totalPage int) {
	intCode := errCode.(int)
	g.c.JSON(httpCode, ResponsePage{
		Code:      intCode,
		Message:   "",
		Data:      data,
		Total:     total,
		TotalPage: totalPage,
	})
}
