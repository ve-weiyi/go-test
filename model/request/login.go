package request

type UserLoginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"`
}
