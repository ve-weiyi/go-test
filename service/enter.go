package service

type serviceGroup struct {
	LoginService LoginService
}

var ServiceStruct = new(serviceGroup)
