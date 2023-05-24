package repo

import (
	"errors"
	"go-app/app/code"
	"go-app/app/http/binding"
	"go-app/app/http/resource"
	"go-app/boot/db"
	"go-app/config"
	"go-app/lib/logger"
	"go-app/lib/rbac"
	"go-app/model"
	"go-app/model/scope"

	"github.com/wumansgy/goEncrypt/hash"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Users = new(usersRepo)

type usersRepo struct{}

func (usersRepo) List(keyword binding.Keyword, pager binding.Pager) (data resource.ApiList, e code.Error) {
	var users []model.UserApi
	query := db.Conn.Model(&model.User{})
	if keyword.Keyword != "" {
		query.Where("username = ? OR phone = ?", keyword.Keyword, keyword.Keyword)
	}
	query.Session(&gorm.Session{})

	// 查询总数用于分页
	var count int64
	query.Count(&count)
	data.Total = count

	query.Scopes(scope.Paginate(pager)).Order("id DESC").Find(&users)
	data.List = users
	return
}

func (usersRepo) Edit(params model.User) (any, e code.Error) {
	var count int64
	if params.ID > 0 {
		db.Conn.Model(&model.User{}).Where("id <> ? AND username = ?", params.ID, params.Username).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "用户名不能重复")
			return
		}
		db.Conn.Model(&model.User{}).Where("id <> ? AND phone = ?", params.ID, params.Phone).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "手机号不能重复")
			return
		}
	} else {
		db.Conn.Model(&model.User{}).Where("username = ?", params.Username).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "用户名不能重复")
			return
		}
		db.Conn.Model(&model.User{}).Where("phone = ?", params.Phone).Count(&count)
		if count > 0 {
			e = code.NewError(-1, "手机号不能重复")
			return
		}
	}

	if params.Password != "" {
		params.Password = hash.HmacSha256Hex([]byte(config.HASH.HmacSha256Key), params.Password)
		err := db.Conn.Save(&params).Error
		if err != nil {
			e = code.NewError(-1, err.Error())
			return
		}
	} else {
		err := db.Conn.Omit("password").Save(&params).Error
		if err != nil {
			e = code.NewError(-1, err.Error())
			return
		}
	}

	return
}

func (usersRepo) Del(id uint) (any, e code.Error) {
	var user model.User
	if db.Conn.First(&user, id).Error != nil {
		e = code.ErrServer
		return
	}
	if err := db.Conn.Delete(&user).Error; err != nil {
		e = code.NewError(-1, err.Error())
	}
	return
}

func (usersRepo) Assign(params binding.UserAssign) (any, e code.Error) {
	var user model.User
	if err := db.Conn.First(&user, params.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = code.NewError(-1, "用户不存在")
			return
		}
		e = code.ErrServer
		return
	}

	if user.Username == "测试用户" || user.Username == "废墟" {
		e = code.NewError(-1, "默认用户不能分配角色")
		return
	}

	_, err := rbac.New().DeleteRolesForUser(user.Username)
	if err != nil {
		logger.Error("[CASBIN]", zap.NamedError("删除用户角色出错", err))
		e = code.ErrServer
		return
	}

	if len(params.Roles) > 0 {
		var roleCodes []string
		db.Conn.Model(&model.Role{}).Where("code IN ?", params.Roles).Pluck("code", &roleCodes)
		_, err := rbac.New().AddRolesForUser(user.Username, roleCodes)
		if err != nil {
			logger.Error("[CASBIN]", zap.NamedError("批量添加用户角色出错", err))
			e = code.ErrServer
			return
		}
	}
	return
}
