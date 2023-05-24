package api

import (
	"go-app/app/http/binding"
	"go-app/app/http/repo"
	"go-app/lib/response"

	"github.com/gin-gonic/gin"
)

var AuthController = new(authController)

type authController struct{}

func (authController) Login(c *gin.Context) response.Json {
	params := c.MustGet("UserLoginParams").(binding.UserLoginParams)
	return response.JsonResponse(repo.Auth.Login(params))
}

func (authController) RefreshToken(c *gin.Context) response.Json {
	params := c.MustGet("RefreshTokenParams").(binding.RefreshTokenParams)
	return response.JsonResponse(repo.Auth.RefreshToken(params))
}
