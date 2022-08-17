package config

import (
	"time"

	"github.com/sakirsensoy/genv"
)

// 数据库线程池设置
type poolConf struct {

	// 连接池里面的连接最大空闲时长
	//
	// 一般不做修改
	ConnMaxIdleTime time.Duration

	// 连接池里面的连接最大存活时长
	//
	// 必须要比MySQL服务器设置的wait_timeout小
	//
	// MySQL一般默认是8小时
	//
	// 一般不做修改
	ConnMaxLifetime time.Duration

	// 连接池里最大空闲连接数
	//
	// 必须要比MaxOpenConns小
	MaxIdleConns int

	// 连接池最多同时打开的连接数
	//
	// 应当比数据库服务器的max_connections值要小
	//
	// 一般设置为： 服务器cpu核心数 * 2 + 服务器有效磁盘数
	MaxOpenConns int
}

var Pool = &poolConf{
	ConnMaxIdleTime: 30 * time.Minute,
	ConnMaxLifetime: 2 * time.Hour,
	MaxIdleConns:    genv.Key("DB_MAX_IDLE_CONNS").Default(4).Int(),
	MaxOpenConns:    genv.Key("DB_MAX_OPEN_CONNS").Default(128).Int(),
}
