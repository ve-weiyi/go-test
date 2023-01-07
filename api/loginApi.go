package api

import (
	"github.com/gin-gonic/gin"
	"go-test/model/request"
	"go-test/model/response"
	"go-test/service"
	"log"
	"net/http"
)

type Login struct {
}

func (l *Login) Login(c *gin.Context) {
	var form request.UserLoginReq //定义request数据
	err := c.Bind(&form)          //前端传递数据赋值给request数据
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //返回json错误信息
		return
	}
	//调用service层登录接口
	msg, err := service.Enter.LoginService.Login(&form) //查不到，err!=nil  ，查到了err=nil
	if err != nil {
		log.Printf("login--->")
		response.FailWithMessage(c, err.Error())
		return
	}
	log.Printf("login--->")
	response.OkWithData(c, msg)
}
