package main

import (
	"github.com/kataras/iris/v12"
	"iriscli/conf"
	"iriscli/web/route"
)

func main() {

	// 初始化
	app := iris.New()

	// 日志级别
	app.Logger().SetLevel("debug")

	// 注册模板
	app.RegisterView(iris.HTML("./web/views", ".html").
		Layout("error/error.html").
		Reload(true))

	// 静态文件
	app.HandleDir("/static", "./web/public")

	// 异常页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问页面出错"))
		ctx.ViewLayout("")
		_ = ctx.View("error/error.html")
	})

	// 注册路由
	route.RegisterRoute(app)

	// 启动服务
	err := app.Run(
		iris.Addr("localhost:" + conf.Sysconfig.Port),
		iris.WithCharset("UTF-8"),
		iris.WithOptimizations,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
	if err != nil {
		panic(err.Error())
	}
}