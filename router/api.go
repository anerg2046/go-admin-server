package router

import (
	"go-app/app/code"
	"go-app/app/http/binding"
	"go-app/app/http/controller/api"
	"go-app/app/http/middleware"
	"go-app/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(r *gin.Engine) {

	r.StaticFS("/upload", http.Dir("./upload"))

	v1 := r.Group("/v1")
	{
		// 无需身份认证的接口
		routerNoneAuth(v1)

		// 需身份认证的接口
		v1.Use(middleware.JWTAuth())
		{
			v1.GET("/get_routers", response.JSON(api.MenuController.Router))

			// 需要casbin权限认证
			v1.Use(middleware.CasbinCheck("/v1"))
			{
				routerSystemSetting(v1)
			}
		}
	}

	r.NoRoute(func(c *gin.Context) {
		// 实现内部重定向
		c.JSON(http.StatusNotFound, response.NewJson().Error(code.ErrRoute))
	})
}

// 无需身份认证的接口
func routerNoneAuth(r *gin.RouterGroup) {
	// 测试接口
	r.GET("/", response.JSON(api.TestController.Ip))
	r.POST("/login", binding.User.Login(), binding.User.VerifyTurnstile(), response.JSON(api.AuthController.Login))
	r.POST("/refreshToken", binding.User.RefreshToken(), response.JSON(api.AuthController.RefreshToken))
}

// 系统管理路由
func routerSystemSetting(r *gin.RouterGroup) {
	// 菜单管理
	sysMenu := r.Group("/sys/menus")
	{
		sysMenu.GET("/list", response.JSON(api.MenuController.List))
		sysMenu.POST("/edit", binding.Menu.Edit(), response.JSON(api.MenuController.Edit))
		sysMenu.POST("/del", binding.Common.ID(), response.JSON(api.MenuController.Del))
	}

	// 角色管理
	sysRole := r.Group("/sys/roles")
	{
		sysRole.GET("/list", response.JSON(api.RoleController.List))
		sysRole.POST("/edit", binding.Role.Edit(), response.JSON(api.RoleController.Edit))
		sysRole.POST("/del", binding.Common.ID(), response.JSON(api.RoleController.Del))
		sysRole.GET("/permission", binding.Common.ID(), response.JSON(api.RoleController.Permission))
		sysRole.POST("/assign", binding.Role.Assign(), response.JSON(api.RoleController.Assign))
	}

	// 用户管理
	sysUser := r.Group("/sys/users")
	{
		sysUser.GET("/list", binding.Common.Keyword(), binding.Common.Pager(), response.JSON(api.UsersController.List))
		sysUser.POST("/edit", binding.Users.Edit(), response.JSON(api.UsersController.Edit))
		sysUser.POST("/del", binding.Common.ID(), response.JSON(api.UsersController.Del))
		sysUser.POST("/assign", binding.Users.Assign(), response.JSON(api.UsersController.Assign))
	}
}
