package server

import (
	"fmt"
	"go-app/app/http/middleware"
	"go-app/config"
	"go-app/lib/logger"
	"go-app/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiServer struct {
	engine *gin.Engine
}

func (s *ApiServer) Start() {
	logger.Info("[Gin]", zap.String("status", "Api Server Start"))
	s.engine.Run(fmt.Sprintf(":%d", config.APP.Port))
}

func NewApiServer(engine *gin.Engine) *ApiServer {
	// 注册特有中间件
	engine.Use(middleware.Cros())
	// 注册路由
	router.RegisterApiRouter(engine)

	return &ApiServer{
		engine: engine,
	}
}
