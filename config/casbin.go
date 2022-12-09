package config

import (
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/sakirsensoy/genv"
)

type casbinConf struct {
	Model          casbinmodel.Model
	DriverName     string
	DataSourceName string
}

var CASBIN = &casbinConf{
	DriverName:     genv.Key("CASBIN_DRIVER").Default("postgres").String(),
	DataSourceName: genv.Key("CASBIN_DSN").Default("db_dsn").String(),
}
