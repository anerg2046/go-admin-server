package middleware

import (
	"go-app/app/code"
	"go-app/config"
	"go-app/lib/logger"
	"go-app/lib/response"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 全局错误崩溃捕获并输出接口信息
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 链接中断，客户端中断连接为正常行为，不需要记录堆栈信息
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					errStr := strings.ToLower(se.Error())
					if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			// 获取用户的请求信息
			httpRequest, _ := httputil.DumpRequest(c.Request, false)
			// 链接中断的情况
			if brokenPipe {
				logger.Error(c.Request.URL.Path,
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				)
				c.Error(err.(error))
				c.Abort()
				// 链接已断开，无法写状态码
				return
			}

			// 如果不是链接中断，就开始记录堆栈信息
			logger.Error("[Recovery from panic]",
				zap.Any("error", err),                      // 记录错误信息
				zap.String("request", string(httpRequest)), // 请求信息
				zap.Stack("stacktrace"),                    // 调用堆栈信息
			)

			if config.APP.Mode == config.MODE_API {
				c.JSON(http.StatusInternalServerError, response.NewJson().Error(code.ErrServer).WithData(errorToString(err)))
			} else if config.APP.Mode == config.MODE_WEB {
				c.Redirect(http.StatusTemporaryRedirect, "/error")
			} else {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
			c.Abort()
			return
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
