package response

import (
	"go-app/app/code"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Json interface {
	i()
	OK() Json
	WithData(data interface{}) Json
	Error(code code.Error) Json
}

type jsonRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// 输出Json数据
// 	args 应该是 [data], [data,err]或者空
// 	err 必须是 code.Error
func JsonResponse(args ...interface{}) Json {
	rsp := NewJson()
	if len(args) == 1 {
		return rsp.OK().WithData(args[0])
	} else if len(args) == 2 {
		if args[0] != nil {
			rsp.WithData(args[0])
		}
		if args[1] != nil {
			return rsp.Error(args[1].(code.Error))
		}
	}
	return rsp.OK()
}

func NewJson() Json {
	return &jsonRsp{}
}

func (r *jsonRsp) i() {}

func (r *jsonRsp) OK() Json {
	r.Code = 0
	r.Msg = "OK"
	return r
}

func (r *jsonRsp) WithData(data interface{}) Json {
	if reflect.ValueOf(data).IsZero() {
		return r
	}
	r.Data = data
	return r
}

func (r *jsonRsp) Error(err code.Error) Json {
	r.Code = err.ErrCode()
	r.Msg = err.ErrMsg()
	return r
}

type jsonHandle func(c *gin.Context) Json

func JSON() func(h jsonHandle) gin.HandlerFunc {
	return func(h jsonHandle) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			rsp := h(ctx)
			ctx.JSON(http.StatusOK, rsp)
		}
	}
}
