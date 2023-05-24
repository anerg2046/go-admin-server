package resource

import (
	"go-app/model"
	"time"
)

type Token struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expires      time.Time `json:"expires"`
}

type Login struct {
	Username string       `json:"username"`
	Roles    model.StrArr `json:"roles"`
	Token
}
