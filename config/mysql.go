package config

import (
	"go-app/lib/util"

	"github.com/sakirsensoy/genv"
)

type mysqlConf struct {
	DSN string
}

var DB = &mysqlConf{
	DSN: util.MysqlDSN(
		genv.Key("DB_HOST").Default("127.0.0.1").String(),
		genv.Key("DB_USER").Default("root").String(),
		genv.Key("DB_PASS").Default("root").String(),
		genv.Key("DB_PORT").Default("3306").String(),
		genv.Key("DB_DBNAME").Default("my_data").String(),
	),
}
