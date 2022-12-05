package request

type UserAddReq struct {
	Username string `json:"username" binding:"required" validate:"required,min=2,max=20"`
	Password string `json:"password" binding:"required"`
}

type UserQueryInfoReq struct {
	Id int64 `form:"id" binding:"required"`
}

type UserUpdateReq struct {
	Id          int64  `json:"id" binding:"required"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	Email       string `json:"email"`
	Picture     string `json:"picture"`
	Phone       string `json:"phone"`
}

type UserRegisterAndLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
