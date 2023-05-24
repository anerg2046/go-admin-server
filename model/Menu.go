package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"go-app/app/enums"
	"strings"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

type Menu struct {
	ID        uint            `form:"id" json:"id,omitempty"`
	ParentID  uint            `form:"parent_id" json:"parent_id"`
	Genre     enums.MenuGenre `form:"genre" json:"genre" binding:"required"`
	Path      string          `form:"path" json:"path,omitempty" binding:"required"`                         // 路由地址
	Component string          `form:"component" json:"component,omitempty"`                                  // 按需加载需要展示的页面
	Name      string          `gorm:"unique;<-:create" form:"name" json:"name,omitempty" binding:"required"` // 路由名字（必须保持唯一）
	Queue     int             `form:"queue" json:"queue"`                                                    // 菜单排序
	Meta      MenuMeta        `gorm:"type:json" form:"meta" json:"meta,omitempty"`                           // 路由元信息
	Children  []Menu          `gorm:"foreignkey:parent_id" json:"children,omitempty"`                        // 子路由
}

func (m *Menu) AfterFind(tx *gorm.DB) (err error) {
	m.Meta.Transition.EnterTransition = "animate__fadeIn animate__faster"
	m.Meta.Transition.LeaveTransition = "animate__fadeOut animate__faster"
	return
}

func (m *Menu) BeforeDelete(tx *gorm.DB) (err error) {
	if strings.Contains(m.Path, "/sys") {
		return errors.New("系统菜单禁止删除")
	}
	return
}

func (m *Menu) BeforeUpdate(tx *gorm.DB) (err error) {
	if strings.Contains(m.Path, "/sys") {
		return errors.New("系统菜单禁止编辑")
	}
	return
}

// 路由元信息
type MenuMeta struct {
	Title        string             `form:"title" json:"title,omitempty"`       // 菜单名称
	Icon         string             `form:"icon" json:"icon,omitempty"`         // 菜单图标
	ShowLink     bool               `form:"showLink" json:"showLink"`           // 是否在菜单中显示
	ShowParent   bool               `form:"showParent" json:"showParent"`       // 是否显示父级菜单
	KeepAlive    bool               `form:"keepAlive" json:"keepAlive"`         // 是否缓存该路由页面（开启后，会保存该页面的整体状态，刷新后会清空状态）
	FrameSrc     string             `form:"frameSrc" json:"frameSrc,omitempty"` // 需要内嵌的iframe链接地址
	FrameLoading bool               `form:"frameLoading" json:"frameLoading"`   // 内嵌的iframe页面是否开启首次加载动画
	HiddenTag    bool               `json:"hiddenTag"`                          // 当前菜单名称或自定义信息禁止添加到标签页
	DynamicLevel int                `json:"dynamicLevel,omitempty"`             // 显示在标签页的最大数量，需满足后面的条件：不显示在菜单中的路由并且是通过query或params传参模式打开的页面
	Transition   MenuMetaTransition `json:"transition,omitempty"`               // 页面加载动画
	// Roles        StrArr             `form:"roles" json:"roles,omitempty"`               // 页面级别权限设置
	// Auths        StrArr             `form:"auths" json:"auths,omitempty"`               // 按钮级别权限设置
}

// 页面加载动画
type MenuMetaTransition struct {
	EnterTransition string `json:"enterTransition,omitempty"` // 当前页面进场动画
	LeaveTransition string `json:"leaveTransition,omitempty"` // 当前页面离场动画
}

func (c MenuMeta) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
func (c *MenuMeta) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, c)
}

func (c MenuMetaTransition) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
func (c *MenuMetaTransition) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, c)
}
