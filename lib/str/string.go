package str

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

func Random(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func VerifyMobileFormat(mobileNum string) bool {
	regular := `^1[3|4|5|6|7|8|9]\d{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 替换括号为中文括号
func CnBrackets(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "(", "（")
	content = strings.ReplaceAll(content, ")", "）")
	return content
}

// 替换括号为英文括号
func EnBrackets(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "（", "(")
	content = strings.ReplaceAll(content, "）", ")")
	return content
}
