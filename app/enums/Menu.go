package enums

type MenuGenre uint16 // 菜单类型

const (
	_           MenuGenre = iota
	Menu_MENU             //菜单
	Menu_ACTION           //操作
	Menu_LINK             //外链
)
