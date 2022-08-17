package request

type ID struct {
	ID uint `form:"id" json:"id" binding:"required,gte=1" label:"ID"`
}

type Genre struct {
	Genre uint16 `form:"genre" json:"genre" binding:"required,gte=1" label:"类型"`
}

type Pager struct {
	Page int `form:"page" json:"page,omitempty" binding:"omitempty,gte=1" label:"页码"`
	Size int `form:"size" json:"size,omitempty" binding:"omitempty,gte=1" label:"每页数量"`
}
