package repo

import (
	"errors"
	"go-app/app/code"
	"go-app/app/enums"
	"go-app/boot/db"
	"go-app/config"
	"go-app/lib/rbac"
	"go-app/lib/util"
	"go-app/model"
	"go-app/model/scope"

	"gorm.io/gorm"
)

var Menu = new(menuRepo)

type menuRepo struct{}

// 用户可显示的菜单
func (menuRepo) Router(auth *config.JwtClaims) (result []model.Menu, e code.Error) {
	var allowMenus []model.Menu
	var menus []model.Menu
	db.Conn.Model(&model.Menu{}).Where("genre IN ?", []enums.MenuGenre{enums.Menu_MENU, enums.Menu_LINK}).Order("queue ASC").Find(&menus)
	for _, menu := range menus {
		if ok, _ := rbac.New().Enforce(auth.Username, menu.Name, "menu"); ok {
			allowMenus = append(allowMenus, menu)
		}
	}
	result = MenuListToTree(allowMenus, model.Menu{})
	return
}

// 菜单列表转树状结构
func MenuListToTree(rows []model.Menu, parent model.Menu) []model.Menu {
	tree := make([]model.Menu, 0)
	for _, row := range rows {
		if row.ParentID == parent.ID {
			row.Children = MenuListToTree(rows, row)
			tree = append(tree, row)
		}
	}
	return tree
}

// 所有菜单列表
func (menuRepo) List() (result []model.Menu, e code.Error) {
	var menus []model.Menu
	db.Conn.Model(&model.Menu{}).Order("queue ASC").Find(&menus)
	result = MenuListToTree(menus, model.Menu{})
	return
}

func (m menuRepo) Edit(params model.Menu) (any, e code.Error) {
	if params.ID > 0 {
		var menu model.Menu
		err := db.Conn.Preload("Children", scope.ExpandChildren).Where("id = ?", params.ID).First(&menu).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				e = code.NewError(-1, "没有这个菜单")
				return
			}
			e = code.ErrServer
			return
		}
		childIDs := m.getAllChildrenID(menu)
		if util.Contain(params.ParentID, childIDs) || params.ParentID == params.ID {
			e = code.NewError(-1, "层级设置错误")
			return
		}
	}
	if params.Genre == enums.Menu_ACTION {
		params.Component = ""
		params.Meta.ShowLink = false
		params.Meta.Icon = ""
	} else if params.Genre == enums.Menu_LINK {
		params.Component = ""
		params.Meta.Icon = ""
	}
	if err := db.Conn.Select("*").Save(&params).Error; err != nil {
		e = code.NewError(-1, err.Error())
	}
	return
}

func (m menuRepo) Del(id uint) (any, e code.Error) {
	var menu model.Menu
	if db.Conn.Preload("Children").First(&menu, id).Error != nil {
		e = code.ErrServer
		return
	}
	if len(menu.Children) > 0 {
		e = code.NewError(-1, "请先删除菜单的下级")
		return
	}
	if err := db.Conn.Delete(&menu).Error; err != nil {
		e = code.NewError(-1, err.Error())
	}
	return
}

func (m menuRepo) getAllChildrenID(menu model.Menu) (ids []uint) {
	for _, child := range menu.Children {
		ids = append(ids, child.ID)
		ids = append(ids, m.getAllChildrenID(child)...)
	}
	return ids
}
