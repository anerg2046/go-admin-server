package router

import (
	"go-app/app/code"
	"go-app/app/http/controller/api"
	"go-app/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(r *gin.Engine) {

	r.StaticFS("/upload", http.Dir("./upload"))

	v1 := r.Group("/v1")
	{
		// 测试接口
		v1.GET("/", response.JSON(api.TestController.Ip))
	}

	r.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.JSON(http.StatusNotFound, response.NewJson().Error(code.ErrRoute))
	})
}
