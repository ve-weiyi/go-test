package service

type serviceGroup struct {
	LoginService    LoginService
	UserInfoService UserInfoService
}

var ServiceStruct = new(serviceGroup)
