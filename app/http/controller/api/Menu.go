package api

import (
	"go-app/app/http/binding"
	"go-app/app/http/repo"
	"go-app/config"
	"go-app/lib/response"
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var MenuController = new(menuController)

type menuController struct{}

func (menuController) Router(c *gin.Context) response.Json {
	auth := c.MustGet("JwtAuth").(*config.JwtClaims)
	return response.JsonResponse(repo.Menu.Router(auth))
}

/** 管理相关接口 */

// 所有菜单
func (menuController) List(c *gin.Context) response.Json {
	return response.JsonResponse(repo.Menu.List())
}

// 添加/编辑菜单
func (menuController) Edit(c *gin.Context) response.Json {
	params := c.MustGet("MenuEditParams").(model.Menu)
	return response.JsonResponse(repo.Menu.Edit(params))
}

// 删除菜单
func (menuController) Del(c *gin.Context) response.Json {
	params := c.MustGet("IdParams").(binding.ID)
	return response.JsonResponse(repo.Menu.Del(params.ID))
}
