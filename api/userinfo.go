package api

import (
	"github.com/gin-gonic/gin"
	"go-test/model/request"
	"go-test/model/response"
	"go-test/service"
	"log"
)

type UserInfo struct {
}

func (u UserInfo) Userinfo(c *gin.Context) {
	log.Printf("Login--->")

	email := c.Query("email")

	var req request.Userinfo

	req.Email = email

	login, err := service.ServiceStruct.UserInfoService.GetUserInfo(&req)

	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	response.OkWithData(c, login)
}
