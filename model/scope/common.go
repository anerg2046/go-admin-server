package scope

import (
	"go-app/app/http/binding"

	"gorm.io/gorm"
)

func Paginate(params binding.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.Size
		return db.Offset(offset).Limit(params.Size)
	}
}

// 展开所有的子集
func ExpandChildren(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", ExpandChildren)
}
