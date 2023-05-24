package binding

import (
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var Role = new(roleBinding)

type roleBinding struct{}

func (roleBinding) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Role
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("RoleEditParams", params)
		c.Next()
	}
}

type RoleAssign struct {
	ID      uint   `form:"id" json:"id" binding:"required" label:"角色ID"`
	MenuIDs []uint `json:"menu_ids,omitempty"`
}

func (roleBinding) Assign() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoleAssign
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("RoleAssignParams", params)
		c.Next()
	}
}
