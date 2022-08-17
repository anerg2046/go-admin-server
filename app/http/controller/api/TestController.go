package api

import (
	"go-app/lib/response"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (TestController) Ip(c *gin.Context) response.Json {
	return response.NewJson().WithData(c.ClientIP())
}
