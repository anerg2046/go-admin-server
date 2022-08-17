package code

var (
	ErrServer         = NewError(10001, "服务异常")       //服务异常
	ErrParam          = NewError(10002, "参数有误")       //参数有误
	ErrToken          = NewError(10003, "Token无效")    //Token无效
	ErrEmptyToken     = NewError(10004, "请求未携带Token") //请求未携带Token
	ErrRoute          = NewError(10005, "请求地址错误")     //请求地址错误
	ErrRecordNotFound = NewError(20001, "未查询到记录")     //未查询到记录

	ErrBusinessTakeLimit = NewError(40001, "商机领取到达上限") //商机领取到达上限
	ErrBusinessTaked     = NewError(40002, "商机已被他人领取") //商机已被他人领取
)
