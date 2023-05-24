package main

import (
	// 必须在import第一位，否则无法获取正确的配置
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
	"go.uber.org/zap"

	"go-app/app/server"
	"go-app/asset"
	"go-app/config"
	"go-app/lib/embedfs"
	"go-app/lib/logger"
	"go-app/lib/validator"
	"html/template"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	config.APP.Mode = config.MODE_WEB
	validator.NewValidator()

	engin := server.NewGinEngine()

	engin.Use(static.Serve("/", embedfs.EmbedFolder(asset.PublicFS, "public", false)))
	engin.StaticFS("/upload", gin.Dir("./upload", false))

	engin.SetHTMLTemplate(parseAssetTemplates())

	app := server.NewWebServer(engin)
	app.Start()

}

// 模板公共方法
func templateFuncs() template.FuncMap {
	return template.FuncMap{}
}

// 从embed读取静态资源
func assetReader(path string, templ *template.Template) {
	fs, _ := asset.ViewFS.ReadDir(path)
	for _, f := range fs {
		if f.IsDir() {
			assetReader(path+"/"+f.Name(), templ)
		} else {
			content, _ := asset.ViewFS.ReadFile(path + "/" + f.Name())
			_, err := templ.Parse(string(content))
			if err != nil {
				logger.Error("[EMBED VIEW]", zap.Error(err))
			}
		}
	}
}

// 从embed读取模板文件
func parseAssetTemplates() *template.Template {
	templ := template.New("").Funcs(templateFuncs())
	assetReader("views", templ)
	return templ
}
