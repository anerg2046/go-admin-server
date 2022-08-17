package web

import (
	"go-app/lib/response"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (TestController) Ip(c *gin.Context) response.Html {
	return response.HtmlResponse("index.html").WithData("ip", c.ClientIP())
}
