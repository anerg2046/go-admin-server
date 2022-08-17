package main

import (
	// 必须在import第一位，否则无法获取正确的配置
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"go-app/app/server"
	"go-app/lib/validator"
)

func main() {
	validator.NewValidator()
	app := server.NewApiServer(server.NewGinEngine())
	app.Start()
}
