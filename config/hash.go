package config

import "github.com/sakirsensoy/genv"

type hashConf struct {
	HmacSha256Key string
	AesKey        string
	AesIV         string
}

var HASH = &hashConf{
	HmacSha256Key: genv.Key("HMAC_SHA256_KEY").Default("MzU0OTkxNzk0NDI1").String(),
	AesKey:        genv.Key("HASH_AES_KEY").Default("6EjDpd8QPnkKXWVC").String(),
	AesIV:         genv.Key("HASH_AES_IV").Default("h8JLYxUk9svkkMad").String(),
}
