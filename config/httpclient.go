package config

import "github.com/sakirsensoy/genv"

type httpclientConf struct {
	Debug      bool
	RetryCount int
	Timeout    int
}

var HTTPCLIENT = &httpclientConf{
	Debug:      genv.Key("HTTP_DEBUG").Default(false).Bool(),
	RetryCount: genv.Key("HTTP_RETRY_COUNT").Default(3).Int(),
	Timeout:    genv.Key("HTTP_TIMEOUT").Default(30).Int(),
}
