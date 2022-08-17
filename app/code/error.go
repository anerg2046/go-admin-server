package code

type Error interface {
	// i 为了避免被其他包实现
	i()
	ErrCode() int
	ErrMsg() string
}

type err struct {
	Code int    // 业务编码
	Msg  string // 错误描述
}

func NewError(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
	}
}

func (e *err) i() {}

func (e *err) ErrCode() int {
	return e.Code
}
func (e *err) ErrMsg() string {
	return e.Msg
}
