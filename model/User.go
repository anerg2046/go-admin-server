package model

import (
	"errors"
	"go-app/lib/rbac"
	"time"

	"gorm.io/gorm"
)

type UserApi struct {
	ID        uint      `form:"id" json:"id,omitempty"`
	Username  string    `form:"username" gorm:"uniqueIndex" json:"username,omitempty"`
	Phone     string    `form:"phone" gorm:"uniqueIndex" json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Roles     StrArr    `gorm:"-" json:"roles"`
}

func (u *UserApi) AfterFind(tx *gorm.DB) (err error) {
	u.Roles, _ = rbac.New().GetRolesForUser(u.Username)
	return
}

type User struct {
	UserApi
	Password string `form:"password" json:"password,omitempty"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.Roles, _ = rbac.New().GetRolesForUser(u.Username)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Username == "admin" || u.Username == "test" {
		return errors.New("默认用户禁止修改")
	}
	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	if u.Username == "admin" || u.Username == "test" {
		return errors.New("默认用户禁止删除")
	}
	rbac.New().DeleteRolesForUser(u.Username)
	return
}

type UserToken struct {
	ID        uint      `json:"id,omitempty"`
	Token     string    `gorm:"index" json:"token,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	User      *User     `gorm:"foreignKey:id" json:"user"`
}
