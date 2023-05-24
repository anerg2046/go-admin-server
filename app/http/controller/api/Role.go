package api

import (
	"go-app/app/http/binding"
	"go-app/app/http/repo"
	"go-app/lib/response"
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var RoleController = new(roleController)

type roleController struct{}

/** 管理相关接口 */

// 所有角色
func (roleController) List(c *gin.Context) response.Json {
	return response.JsonResponse(repo.Role.List())
}

// 添加/编辑角色
func (roleController) Edit(c *gin.Context) response.Json {
	params := c.MustGet("RoleEditParams").(model.Role)
	return response.JsonResponse(repo.Role.Edit(params))
}

// 删除角色
func (roleController) Del(c *gin.Context) response.Json {
	params := c.MustGet("IdParams").(binding.ID)
	return response.JsonResponse(repo.Role.Del(params.ID))
}

// 角色权限
func (roleController) Permission(c *gin.Context) response.Json {
	params := c.MustGet("IdParams").(binding.ID)
	return response.JsonResponse(repo.Role.Permission(params.ID))
}

// 角色指派权限
func (roleController) Assign(c *gin.Context) response.Json {
	params := c.MustGet("RoleAssignParams").(binding.RoleAssign)
	return response.JsonResponse(repo.Role.Assign(params))
}
