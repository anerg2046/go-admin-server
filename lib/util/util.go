package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 检查obj是否包含在target里
func Contain[T comparable](obj T, target []T) bool {
	for _, v := range target {
		if obj == v {
			return true
		}
	}
	return false
}

// Pretty 友好显示控制台输出数据
func Pretty(data any) {
	src, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(src))
}

// 判断指针是否为空
func IsNil(i any) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

// 返回任意数据的指针
func Ptr(i any) *any {
	return &i
}
