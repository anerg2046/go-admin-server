package server

import (
	"go-app/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewGinEngine() (engine *gin.Engine) {
	// 一般来说不需要开启gin的debug模式
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	// 注册全局异常处理
	engine.Use(middleware.Recover)
	return
}
