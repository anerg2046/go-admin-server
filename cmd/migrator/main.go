package main

import (
	// 必须在import第一位，否则无法获取正确的配置
	"time"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"
	"gorm.io/gorm"

	"go-app/app/enums"
	"go-app/boot/db"
	"go-app/config"
	"go-app/lib/database"
	"go-app/lib/rbac"
	"go-app/model"
)

func main() {
	conn, _ := database.ConnDB(config.DB_APP.DSN, config.DB_APP.DB_TYPE)
	app(conn)
	defaultMenu(conn)
	defaultRole(conn)
	defaultUser(conn)
}

func app(conn *gorm.DB) {
	conn.AutoMigrate(
		&model.User{},
		&model.UserToken{},
		&model.Menu{},
		&model.Role{},
	)
}

func defaultMenu(conn *gorm.DB) {
	var count int64
	db.Conn.Model(&model.Menu{}).Count(&count)
	if count > 0 {
		return
	}
	menu := []model.Menu{
		{
			Genre: enums.Menu_MENU,
			Path:  "/sys",
			Name:  "SystemSetting",
			Queue: 0,
			Meta: model.MenuMeta{
				Title:    "系统管理",
				Icon:     "solar:settings-linear",
				ShowLink: true,
			},
			Children: []model.Menu{
				{
					Genre: enums.Menu_MENU,
					Path:  "/sys/menu/index",
					Name:  "SysMenuManagement",
					Queue: 0,
					Meta: model.MenuMeta{
						Title:      "菜单管理",
						KeepAlive:  true,
						ShowParent: true,
						ShowLink:   true,
					},
					Children: []model.Menu{
						{
							Genre: enums.Menu_ACTION,
							Path:  "/sys/menus/list",
							Name:  "ActMenuList",
							Meta: model.MenuMeta{
								Title:      "菜单-列表",
								KeepAlive:  false,
								ShowParent: true,
								ShowLink:   false,
							},
						},
					},
				},
				{
					Genre:     enums.Menu_MENU,
					Path:      "/sys/roles",
					Component: "/sys/role/index",
					Name:      "SysRoleManagement",
					Queue:     5,
					Meta: model.MenuMeta{
						Title:      "角色管理",
						KeepAlive:  true,
						ShowParent: true,
						ShowLink:   true,
					},
				},
				{
					Genre:     enums.Menu_MENU,
					Path:      "/sys/users",
					Component: "/sys/user/index",
					Name:      "SysUserManagement",
					Queue:     10,
					Meta: model.MenuMeta{
						Title:      "用户管理",
						KeepAlive:  true,
						ShowParent: true,
						ShowLink:   true,
					},
					Children: []model.Menu{
						{
							Genre: enums.Menu_ACTION,
							Path:  "/sys/users/list",
							Name:  "ActUserList",
							Meta: model.MenuMeta{
								Title:      "用户-列表",
								KeepAlive:  false,
								ShowParent: true,
								ShowLink:   false,
							},
						},
					},
				},
			},
		},
	}
	conn.Save(&menu)
}

func defaultRole(conn *gorm.DB) {
	var count int64
	conn.Model(&model.Role{}).Where(&model.Role{Code: "superadmin"}).Count(&count)
	if count == 0 {
		role := []model.Role{
			{
				Code:        "superadmin",
				Name:        "超级管理员",
				Queue:       0,
				Description: "拥有所有权限，且不可删除",
			},
			{
				Code:        "user",
				Name:        "注册用户",
				Queue:       5,
				Description: "普通用户，只拥有普通权限",
			},
		}
		conn.Create(&role)
	}
}

func defaultUser(conn *gorm.DB) {
	var count int64
	db.Conn.Model(&model.User{}).Count(&count)
	if count > 0 {
		return
	}

	users := []model.User{
		{
			UserApi: model.UserApi{
				Username:  "admin",
				Phone:     "13000000000",
				CreatedAt: time.Now(),
			},
			Password: "admin1234",
		},
		{
			UserApi: model.UserApi{
				Username:  "test",
				Phone:     "13100000000",
				CreatedAt: time.Now(),
			},
			Password: "test123456",
		},
	}

	conn.Create(&users)

	rbac.New().AddRoleForUser("admin", "superadmin")
	rbac.New().AddRoleForUser("test", "user")
}
