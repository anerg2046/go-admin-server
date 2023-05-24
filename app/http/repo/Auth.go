package repo

import (
	"errors"
	"go-app/app/code"
	"go-app/app/http/binding"
	"go-app/app/http/resource"
	"go-app/boot/db"
	"go-app/config"
	"go-app/model"
	"time"

	jwtlib "go-app/lib/jwt"
	"go-app/lib/str"

	"github.com/golang-jwt/jwt"
	"github.com/wumansgy/goEncrypt/hash"
	"gorm.io/gorm"
)

var Auth = new(authRepo)

type authRepo struct{}

func (ar *authRepo) Login(params binding.UserLoginParams) (result resource.Login, e code.Error) {
	var user model.User
	err := db.Conn.Where("LOWER(username) = LOWER(?) OR phone = ?", params.Username, params.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = code.NewError(-1, "用户不存在")
			return
		}
		e = code.ErrServer
		return
	}

	if user.Password != hash.HmacSha256Hex([]byte(config.HASH.HmacSha256Key), params.Password) {
		e = code.NewError(-1, "密码不正确")
		return
	}

	result.Username = user.Username
	result.AccessToken, result.Expires = ar.genJwtToken(user)
	result.RefreshToken = ar.genRefreshToken(user)
	result.Roles = user.Roles
	return
}

func (ar *authRepo) RefreshToken(params binding.RefreshTokenParams) (result resource.Token, e code.Error) {
	var userToken model.UserToken
	err := db.Conn.Preload("User").Where(&model.UserToken{Token: params.RefreshToken}).First(&userToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e = code.ErrToken
			return
		}
		e = code.ErrServer
		return
	}

	if userToken.ExpiredAt.Before(time.Now()) {
		e = code.ErrToken
		return
	}

	result.AccessToken, result.Expires = ar.genJwtToken(*userToken.User)
	result.RefreshToken = userToken.Token
	return
}

// 生成jwt token即AccessToken
func (authRepo) genJwtToken(user model.User) (string, time.Time) {
	exp := time.Now().Add(time.Minute * 10)
	claims := config.JwtClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
	}

	token, err := jwtlib.CreateToken(claims)
	if err != nil {
		panic(err)
	}
	return token, exp
}

// 生成随机数token即RefreshToken
func (authRepo) genRefreshToken(user model.User) string {
	var userToken model.UserToken
	db.Conn.Where(&model.UserToken{ID: user.ID}).FirstOrInit(&userToken)
	userToken.Token = str.Random(128)
	userToken.ExpiredAt = time.Now().AddDate(0, 0, 30)
	db.Conn.Save(&userToken)

	return userToken.Token
}
