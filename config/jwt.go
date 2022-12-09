package config

import (
	"github.com/sakirsensoy/genv"

	"github.com/golang-jwt/jwt"
)

type jwtConf struct {
	Key         string
	HeaderField string
	RedirectUrl string
}

// 载荷，可以加一些自己需要的信息
type JwtClaims struct {
	ID     uint   `json:"id,omitempty"`
	Nick   string `json:"nick,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	jwt.StandardClaims
}

var JWT = &jwtConf{
	Key:         genv.Key("JWT_KEY").Default("5fqwcChUsM*g@F2G").String(),
	HeaderField: genv.Key("JWT_HEADER_FIELD").Default("Authorization").String(),
	RedirectUrl: genv.Key("JWT_REDIRECT_URL").Default("/login").String(),
}
