package main

import (
	// 必须在import第一位，否则无法获取正确的配置
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"go-app/app/server"
	"go-app/config"
	"go-app/lib/validator"
)

func main() {
	config.APP.Mode = config.MODE_API
	validator.NewValidator()
	app := server.NewApiServer(server.NewGinEngine())
	app.Start()
}
