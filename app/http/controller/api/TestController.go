package api

import (
	"go-app/lib/response"

	"github.com/gin-gonic/gin"
)

var TestController = new(testController)

type testController struct{}

func (testController) Ip(c *gin.Context) response.Json {
	return response.NewJson().WithData(c.ClientIP())
}
