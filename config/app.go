package config

import (
	"time"

	"github.com/sakirsensoy/genv"
)

type APP_MODE uint16

const (
	_ APP_MODE = iota
	MODE_WEB
	MODE_API
)

type appConf struct {
	Debug    bool
	Env      string
	Mode     APP_MODE // 运行模式 web 或者 api
	Timezone *time.Location
	Port     int
}

var APP *appConf

func init() {
	APP = &appConf{
		Debug:    genv.Key("APP_DEBUG").Default(false).Bool(),
		Timezone: getTimezone(),
		Port:     genv.Key("APP_PORT").Default(8010).Int(),
	}
}

func getTimezone() (timezone *time.Location) {
	timezone, _ = time.LoadLocation(genv.Key("APP_TIMEZONE").Default("PRC").String())
	return
}
