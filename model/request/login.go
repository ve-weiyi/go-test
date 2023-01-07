package request

type UserLoginReq struct {
	Uloginid  string `form:"uloginid"`  // 用户名
	Upassword string `form:"upassword"` //密码
}
