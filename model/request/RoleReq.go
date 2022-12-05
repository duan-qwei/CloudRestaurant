package request

type RoleAddReq struct {
	Name string `json:"name" binding:"required"`
}

type RoleUpdateReq struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
