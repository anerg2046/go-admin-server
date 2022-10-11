package binding

// 一些通用绑定及验证

import (
	"go-app/app/code"
	"go-app/config"
	"go-app/lib/response"
	"go-app/lib/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Common struct{}

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

func (Common) ID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var id ID
		params, err := bindParams(c, id)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("IdParams", params.ID)
		c.Next()
	}
}

func (Common) Genre() gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre Genre
		params, err := bindParams(c, genre)
		if err != nil {
			c.Abort()
			return
		}
		c.Set("GenreParams", params.Genre)
		c.Next()
	}
}

func (Common) Pager() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pager Pager
		params, err := bindParams(c, pager)
		if err != nil {
			c.Abort()
			return
		}
		if params.Page == 0 {
			params.Page = 1
		}
		if params.Size == 0 {
			params.Size = 50
		}
		c.Set("PagerParams", params)
		c.Next()
	}
}

func bindParams[T any](c *gin.Context, params T) (T, error) {
	var err error
	if c.ContentType() == "application/json" {
		err = c.ShouldBindBodyWith(&params, binding.JSON)
	} else {
		err = c.ShouldBind(&params)
	}
	if err != nil {
		if config.APP.Mode == config.MODE_API {
			c.JSON(http.StatusOK, response.NewJson().Error(code.ErrParam).WithData(validator.Error(err)))
		} else if config.APP.Mode == config.MODE_WEB {
			c.Redirect(http.StatusTemporaryRedirect, "/404")
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
	return params, err
}
