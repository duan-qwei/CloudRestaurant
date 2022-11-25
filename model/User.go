package model

type User struct {
	BaseModel
	Password string `json:"password" valid:"Required"`
	Username string `json:"username" valid:"Required"`
}
