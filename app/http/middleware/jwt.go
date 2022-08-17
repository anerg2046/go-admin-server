package middleware

import (
	"go-app/app/code"
	"go-app/config"
	"go-app/lib/jwt"
	"go-app/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.APP.Env == "local" {
			claims := &jwt.CustomClaims{ID: 1, Nick: "Coeus", Avatar: ""}
			c.Set("claims", claims)
		} else {
			token := c.Request.Header.Get("token")
			if token == "" {
				if config.APP.Mode == config.MODE_API {
					c.JSON(http.StatusOK, response.NewJson().Error(code.ErrEmptyToken))
				} else if config.APP.Mode == config.MODE_WEB {
					c.Redirect(http.StatusTemporaryRedirect, "/login")
				} else {
					c.AbortWithStatus(http.StatusForbidden)
				}
				c.Abort()
				return
			}

			// parseToken 解析token包含的信息
			claims, err := jwt.ParseToken(token)
			if err != nil {
				if config.APP.Mode == config.MODE_API {
					c.JSON(http.StatusOK, response.NewJson().Error(code.ErrToken))
				} else if config.APP.Mode == config.MODE_WEB {
					c.Redirect(http.StatusTemporaryRedirect, "/login")
				} else {
					c.AbortWithStatus(http.StatusForbidden)
				}
				c.Abort()
				return
			}
			// 继续交由下一个路由处理,并将解析出的信息传递下去
			c.Set("claims", claims)
		}
	}
}
