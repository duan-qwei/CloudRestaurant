package vo

type SysUser struct {
	Password string `json:"password" valid:"Required"`
	Username string `json:"username" valid:"Required"`
}
