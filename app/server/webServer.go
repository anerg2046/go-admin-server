package server

import (
	"fmt"
	"go-app/config"
	"go-app/lib/logger"
	"go-app/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type webServer struct {
	engine *gin.Engine
}

func (s *webServer) Start() {
	logger.Info("[Gin]", zap.String("status", "Web Server Start"))
	s.engine.Run(fmt.Sprintf(":%d", config.APP.Port))
}

func NewWebServer(engine *gin.Engine) *webServer {
	engine.SetTrustedProxies(nil)
	// 注册路由
	router.RegisterWebRouter(engine)

	return &webServer{
		engine: engine,
	}
}
