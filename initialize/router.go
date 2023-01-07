package initialize

import (
	"github.com/gin-gonic/gin"
	"go-test/api"
	"net/http"
)

func RouterInit() *gin.Engine {
	// 1.创建路由
	r := gin.Default()

	group := r.Group("api/")

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	group.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	group.GET("/login", api.API.LoginAPi.Login)
	return r
}
