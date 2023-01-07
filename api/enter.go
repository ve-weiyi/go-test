package api

type apiGroup struct {
	LoginAPi Login
	UserApi  User
}

var API = new(apiGroup)
