package api

type apiGroup struct {
	LoginAPi    Login
	UserApi     User
	UserInfoApi UserInfo //变量  类型
}

var API = new(apiGroup)
