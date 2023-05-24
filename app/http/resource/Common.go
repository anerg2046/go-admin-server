package resource

// 通用Api列表返回数据
type ApiList struct {
	List  any   `json:"list"`
	Total int64 `json:"total"`
}
