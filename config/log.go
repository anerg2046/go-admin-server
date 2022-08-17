package config

import "github.com/sakirsensoy/genv"

type logConf struct {
	Level     string
	Type      string
	Path      string
	File      string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Compress  bool
}

var LOG = &logConf{
	Level:     genv.Key("LOG_LEVEL").Default("error").String(),
	Type:      genv.Key("LOG_TYPE").Default("stdout").String(),
	Path:      genv.Key("LOG_PATH").Default("storage/logs/").String(),
	File:      genv.Key("LOG_FILE").Default("logs.log").String(),
	MaxSize:   genv.Key("LOG_MAX_SIZE").Default(64).Int(),
	MaxBackup: genv.Key("LOG_MAX_BACKUP").Default(10).Int(),
	MaxAge:    genv.Key("LOG_MAX_AGE").Default(7).Int(),
	Compress:  genv.Key("LOG_COMPRESS").Default(false).Bool(),
}
