package request

type UserAddReq struct {
	Username string `json:"username" validate:"required,min=2,max=20"`
	Password string `json:"password"`
}
