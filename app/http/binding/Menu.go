package binding

import (
	"go-app/model"

	"github.com/gin-gonic/gin"
)

var Menu = new(menuBinding)

type menuBinding struct{}

func (u menuBinding) Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Menu
		params, err := bindParams(c, data)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("MenuEditParams", params)
		c.Next()
	}
}
