package binding

import (
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var Users = new(usersBinding)

type usersBinding struct{}

func (usersBinding) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.User
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("UsersEditParams", params)
		c.Next()
	}
}

type UserAssign struct {
	ID    uint     `form:"id" json:"id" binding:"required" label:"用户ID"`
	Roles []string `form:"roles" json:"roles,omitempty"`
}

func (usersBinding) Assign() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserAssign
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("UserAssignParams", params)
		c.Next()
	}
}
