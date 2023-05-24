package api

import (
	"go-app/app/http/binding"
	"go-app/app/http/repo"
	"go-app/lib/response"
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var UsersController = new(usersController)

type usersController struct{}

/** 管理相关接口 */

// 所有用户
func (usersController) List(c *gin.Context) response.Json {
	keyword := c.MustGet("KeywordParams").(binding.Keyword)
	pager := c.MustGet("PagerParams").(binding.Pager)
	return response.JsonResponse(repo.Users.List(keyword, pager))
}

// 添加/编辑用户
func (usersController) Edit(c *gin.Context) response.Json {
	params := c.MustGet("UsersEditParams").(model.User)
	return response.JsonResponse(repo.Users.Edit(params))
}

// 删除用户
func (usersController) Del(c *gin.Context) response.Json {
	params := c.MustGet("IdParams").(binding.ID)
	return response.JsonResponse(repo.Users.Del(params.ID))
}

// 用户分配角色
func (usersController) Assign(c *gin.Context) response.Json {
	params := c.MustGet("UserAssignParams").(binding.UserAssign)
	return response.JsonResponse(repo.Users.Assign(params))
}
