package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/sessions"
	"github.com/tddey01/go_bas/iris_web/002/002/config"
)

// newApp()  构建
func newApp() *iris.Application {
	app := iris.New()

	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")

	//注册视图文件
	app.RegisterView(iris.HTML("/static", ".html"))
	app.Get("/", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.View("index.html")
	})
	return app
}

// mvcHandler  * MVC 架构模式处理

func mvcHandler(app *iris.Application) {
	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	//engine := datasource.
	fmt.Println(sessManager)
}

func main() {
	// app := newApp()

	// 应用设置
	//configuration(app)

	// 路由设置
	//mvcHandle(app)

	config := config.InitConfig()
	fmt.Println(config)
}
