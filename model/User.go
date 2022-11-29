package model

type User struct {
	Id int64 `uri:"id" binding:"required"`
	BaseModel
	Password string `json:"password" valid:"Required"`
	Username string `json:"username" valid:"Required"`
}
