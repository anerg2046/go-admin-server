package jwt

import (
	"errors"
	"go-app/config"

	"github.com/golang-jwt/jwt"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID     uint   `json:"id,omitempty"`
	Nick   string `json:"nick,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	jwt.StandardClaims
}

// 一些常量
var (
	ErrTokenExpired     error = errors.New("Token已过期")
	ErrTokenNotValidYet error = errors.New("Token还未生效")
	ErrTokenMalformed   error = errors.New("Token无法解析")
	ErrTokenInvalid     error = errors.New("Token无效")
)

type jwtAuth struct {
	SigningKey []byte
}

var jwtauth *jwtAuth

func init() {
	jwtauth = &jwtAuth{
		SigningKey: []byte(config.JWT.Key),
	}
}

func CreateToken(claims CustomClaims) (string, error) {
	return jwtauth.createToken(claims)
}

func (j *jwtAuth) createToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func ParseToken(token string) (*CustomClaims, error) {
	return jwtauth.parseToken(token)
}

func (j *jwtAuth) parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
