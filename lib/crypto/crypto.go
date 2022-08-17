package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"go-app/lib/str"
)

type crypto struct {
	Key []byte
	Iv  []byte
}

func NewCrypto() *crypto {
	return &crypto{
		getKey(),
		getIv(),
	}
}

var Crypto = NewCrypto()

func getKey() []byte {
	return []byte(str.Random(16))
}

func getIv() []byte {
	return []byte(str.Random(16))
}

// aes加密, 分组模式ctr
func (c *crypto) Encrypt(plainText string) string {
	plainText = str.Reverse(plainText)
	// 1. 建一个底层使用aes的密码接口
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr分组接口
	stream := cipher.NewCTR(block, c.Iv)

	// 4. 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, []byte(plainText))
	// fmt.Println(cipherText)

	return fmt.Sprintf("%x", cipherText)
}

// des解密
func (c *crypto) Decrypt(cipherText string) string {
	decode, err := hex.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}
	// 1. 建一个底层使用des的密码接口
	block, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用ctr模式解密的接口
	stream := cipher.NewCTR(block, c.Iv)
	// 3. 解密
	stream.XORKeyStream(decode, decode)

	// fmt.Println(decode)

	return str.Reverse(string(decode))
}
