package config

import "github.com/sakirsensoy/genv"

type jwtConf struct {
	Key string
}

var JWT = &jwtConf{
	Key: genv.Key("JWT_KEY").Default("5fqwcChUsM*g@F2G").String(),
}
