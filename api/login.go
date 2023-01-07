package api

import (
	"github.com/gin-gonic/gin"
	"go-test/model/request"
	"go-test/model/response"
	"go-test/service"
	"log"
)

type Login struct {
}

func (l *Login) Login(c *gin.Context) {

	log.Printf("Login--->")

	name := c.Query("name")
	pwd := c.Query("pwd")

	var user request.UserLoginReq

	user.Username = name
	user.Password = pwd

	login, err := service.ServiceStruct.LoginService.Login(&user)

	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	response.OkWithData(c, login)
}
