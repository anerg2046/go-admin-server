package repo

import (
	"go-app/app/code"
	"go-app/app/enums"
	"go-app/app/http/binding"
	"go-app/boot/db"
	"go-app/lib/logger"
	"go-app/lib/rbac"
	"go-app/model"

	"go.uber.org/zap"
)

var Role = new(roleRepo)

type roleRepo struct{}

func (roleRepo) List() (result []model.Role, e code.Error) {
	db.Conn.Order("queue ASC").Find(&result)
	return
}

func (roleRepo) Edit(params model.Role) (any, e code.Error) {
	var count int64
	if params.ID > 0 {
		db.Conn.Model(&model.Role{}).Where("id <> ? AND code = ?", params.ID, params.Code).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "角色编码不能重复")
			return
		}
	} else {
		db.Conn.Model(&model.Role{}).Where("code = ?", params.Code).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "角色编码不能重复")
			return
		}
	}
	err := db.Conn.Save(&params).Error
	if err != nil {
		e = code.ErrServer
	}
	return
}

func (roleRepo) Del(id uint) (any, e code.Error) {
	var role model.Role
	if db.Conn.First(&role, id).Error != nil {
		e = code.ErrServer
		return
	}
	if err := db.Conn.Delete(&role).Error; err != nil {
		e = code.NewError(-1, err.Error())
	}
	return
}

// 获取角色的权限
func (roleRepo) Permission(id uint) (menuIDs []uint, e code.Error) {
	var role model.Role
	db.Conn.Where("id = ?", id).First(&role)

	var menuNames []string
	permissions := rbac.New().GetPermissionsForUser(role.Code)
	for _, p := range permissions {
		menuNames = append(menuNames, p[1])
	}

	// 必须加上parent_id排序，否则前端显示可能有bug
	// 因为在角色指派权限的时候需要设置已有的，而如果先设置子叶，再设置父叶，会出现半选变成全选的问题
	db.Conn.Model(&model.Menu{}).Where("name IN ? OR path IN ?", menuNames, menuNames).Order("parent_id ASC").Pluck("id", &menuIDs)

	return
}

// 角色指派权限
func (roleRepo) Assign(params binding.RoleAssign) (any, e code.Error) {
	var role model.Role
	db.Conn.Where("id = ?", params.ID).First(&role)

	// 先要删除该角色 所有权限
	rbac.New().DeletePermissionsForUser(role.Code)

	var menus []model.Menu
	db.Conn.Model(&model.Menu{}).Where("id IN ?", params.MenuIDs).Find(&menus)
	for _, menu := range menus {
		if menu.Genre == enums.Menu_ACTION {
			//针对服务端的接口权限
			_, err := rbac.New().AddPermissionForUser(role.Code, menu.Path, "request")
			if err != nil {
				logger.Error("[Casbin]", zap.Error(err))
				e = code.ErrServer
				return
			}
		} else { //菜单显示权限
			_, err := rbac.New().AddPermissionForUser(role.Code, menu.Name, "menu")
			if err != nil {
				logger.Error("[Casbin]", zap.Error(err))
				e = code.ErrServer
				return
			}
		}
	}
	return
}
