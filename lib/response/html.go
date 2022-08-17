package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Html interface {
	i()
	Template(args ...string) string
	WithData(key string, data any) Html
	WithFunc(key string, data func()) Html
	Payload() any
}

type htmlRsp struct {
	Tpl  string
	Data map[string]any
	Func map[string]func()
}

func NewHtml() Html {
	return &htmlRsp{}
}

func (r *htmlRsp) i() {}

func (r *htmlRsp) Template(args ...string) string {
	if len(args) == 1 {
		r.Tpl = args[0]
	}
	return r.Tpl
}
func (r *htmlRsp) Payload() any {
	return r.Data
}

func (r *htmlRsp) WithData(key string, data any) Html {
	if r.Data == nil {
		r.Data = make(map[string]any)
	}
	r.Data[key] = data
	return r
}

func (r *htmlRsp) WithFunc(key string, data func()) Html {
	if r.Func == nil {
		r.Func = make(map[string]func())
	}
	r.Func[key] = data
	return r
}

func HtmlResponse(tpl string) Html {
	rsp := NewHtml()
	rsp.Template(tpl)
	return rsp
}

type htmlHandle func(c *gin.Context) Html

func HTML() func(h htmlHandle) gin.HandlerFunc {
	return func(h htmlHandle) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			rsp := h(ctx)
			ctx.HTML(http.StatusOK, rsp.Template(), rsp.Payload())
		}
	}
}
