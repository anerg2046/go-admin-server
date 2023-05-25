package config

import (
	"go-app/lib/util"

	"github.com/sakirsensoy/genv"
)

type DBTYPE uint16

const (
	_ DBTYPE = iota
	DBTYPE_MYSQL
	DBTYPE_MSSQL
	DBTYPE_POSTGRES
)

type dbConf struct {
	DSN     string
	DB_TYPE DBTYPE
}

var DB_APP = &dbConf{
	DSN: util.PostgresDSN(
		genv.Key("DB_APP_HOST").Default("postgres").String(),
		genv.Key("DB_APP_USER").Default("root").String(),
		genv.Key("DB_APP_PASS").Default("111222").String(),
		genv.Key("DB_APP_PORT").Default("5432").String(),
		genv.Key("DB_APP_DBNAME").Default("go_admin").String(),
	),
	DB_TYPE: DBTYPE_POSTGRES,
}
