package binding

// 一些通用绑定及验证

import (
	"go-app/app/code"
	"go-app/app/http/request"
	"go-app/config"
	"go-app/lib/response"
	"go-app/lib/validator"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Common struct{}

func (Common) ID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var params request.ID
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
			c.Abort()
			return
		}
		c.Set("ID", params.ID)
		c.Next()
	}
}

func (Common) Genre() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params request.Genre
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
			c.Abort()
			return
		}
		c.Set("genre", params.Genre)
		c.Next()
	}
}

func (Common) Pager() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params request.Pager
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
			c.Abort()
			return
		}
		if params.Page == 0 {
			params.Page = 1
		}
		if params.Size == 0 {
			params.Size = 50
		}
		c.Set("pager", params)
		c.Next()
	}
}
