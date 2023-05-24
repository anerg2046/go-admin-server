package model

import (
	"errors"
	"go-app/lib/rbac"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	ID          uint   `form:"id" json:"id" binding:"omitempty,gte=1" label:"角色ID"`       // 角色id
	Code        string `form:"code" json:"code" binding:"required" label:"角色编码"`          // 角色编码
	Name        string `form:"name" json:"name" binding:"required" label:"角色名称"`          // 角色名称
	Queue       uint   `form:"queue" json:"queue" binding:"omitempty,gte=0" label:"角色排序"` // 排序
	Description string `form:"description" json:"description,omitempty"`
}

func (r *Role) BeforeDelete(tx *gorm.DB) (err error) {
	if r.Code == "user" || r.Code == "superadmin" {
		return errors.New("默认角色禁止删除")
	}
	rbac.New().DeleteRole(r.Code)
	rbac.New().DeletePermissionsForUser(r.Code)
	return
}
