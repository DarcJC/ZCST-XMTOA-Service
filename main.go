package main

import (
	_ "Onboarding/docs"
	"Onboarding/handler"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mix-go/dotenv"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo
var address string
var port int

// @title OA入职微服务
// @version 1.0
// @description 入职微服务
// @contact.name DarcJC
// @contact.email i@darc.pro
// @host localhost
// @BasePath /v1
func main() {
	// 运行服务端
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", address, port)))
}

func loadHandlers() {
	handler.UseUserHandler(e)

	// 添加文档页面
	e.GET("/docs/*", echoSwagger.WrapHandler)
}

func init() {
	// 创建Echo服务端实例
	e = echo.New()

	// 如果有.env文件 则加载
	if err := dotenv.Load(".env"); err != nil {
		e.Logger.Info("Could not found .env file, skipped.")
	}

	// 加载Handler
	loadHandlers()

	// 获取绑定的地址 默认为127.0.0.1
	address = dotenv.Getenv("HOST").String("127.0.0.1")
	// 获取绑定的端口号 默认8000
	port = int(dotenv.Getenv("PORT").Int64(8000))
}
