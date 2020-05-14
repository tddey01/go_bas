package main

import (
	"github.com/go_bas/iris_web/002/002/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/go_bas/iris_web/002/002/service"
	"github.com/go_bas/iris_web/002/002/controller"
	"github.com/go_bas/iris_web/002/002/datasource"
	"github.com/kataras/iris/sessions"
	"time"
)

// newApp()  构建
func newApp() *iris.Application {
	app := iris.New()
	//设置日志级别  开发阶段为debug
	app.Logger().SetLevel("debug")

	// 注册静态资源
	app.StaticWeb()
}

func main() {
	app := newApp()

	// 应用设置
	configuration(app)

	// 路由设置
	mvcHandle(app)

	config := config.InitConfig()
}
