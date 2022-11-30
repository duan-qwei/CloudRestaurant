package request

type UserAddReq struct {
	Username string `json:"username" binding:"required" validate:"required,min=2,max=20"`
	Password string `json:"password" binding:"required"`
}

type UserQueryInfoReq struct {
	Id int64 `form:"id" binding:"required"`
}
