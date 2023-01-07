package initialize

import (
	"github.com/gin-gonic/gin"
	"go-test/api"
	middleware "go-test/middelware"
	"net/http"
)

func RouterInit() *gin.Engine {
	// 1.创建路由
	r := gin.Default()
	r.Use(middleware.Cors()) //解决跨域
	group := r.Group("api/") //api分组
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	group.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	//group.POST("user/login", api.API.LoginAPi.Login) //接口API
	group.POST("user/login", api.EnterApi.LoginAPi.Login) //接口API

	//group.POST("/userinfo", api.API.UserInfoApi.Userinfo)
	return r

}
