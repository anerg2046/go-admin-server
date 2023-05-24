package binding

import (
	"go-app/app/code"
	"go-app/lib/httpclient"
	"go-app/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var User = new(userBinding)

type userBinding struct{}

type UserLoginParams struct {
	Username       string `form:"username" binding:"required" label:"帐号"`
	Password       string `form:"password" binding:"required" label:"密码"`
	TurnstileToken string `form:"turnstileToken" binding:"required" label:"TurnstileToken"`
}

func (u userBinding) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserLoginParams
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("UserLoginParams", params)
		c.Set("TurnstileToken", params.TurnstileToken)
		c.Next()
	}
}

func (u userBinding) VerifyTurnstile() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := map[string]string{
			"secret":   "0x4AAAAAAADTj7oP4w2VowOc8GAxonfNv6E",
			"response": c.MustGet("TurnstileToken").(string),
		}
		url := "https://challenges.cloudflare.com/turnstile/v0/siteverify"
		var result = struct {
			Success bool `json:"success,omitempty"`
		}{}
		httpclient.C().R().SetResult(&result).SetFormData(payload).Post(url)
		if !result.Success {
			c.JSON(http.StatusUnauthorized, response.NewJson().Error(code.NewError(-1, "未通过人机验证")))
			c.Abort()
			return
		}
		c.Next()
	}
}

type RefreshTokenParams struct {
	RefreshToken string `form:"refreshToken" binding:"required" label:"Token"`
}

func (userBinding) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RefreshTokenParams
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("RefreshTokenParams", params)
		c.Next()
	}
}
