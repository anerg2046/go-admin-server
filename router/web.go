package router

import (
	"go-app/app/http/controller/web"
	"go-app/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(r *gin.Engine) {

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	TestController := new(web.TestController)
	r.GET("/", response.HTML(TestController.Ip))

	r.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})
}
