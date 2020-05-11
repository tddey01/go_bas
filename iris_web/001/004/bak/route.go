package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/context"
// 路由中间件
func userMiddleware(context iris.Context){
	context.Next()
}
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	//用户模块users
	//  xxx/users/register 注册
	//  xxx/users/login  登录
	//  xxx/users/info   获取用户信息

	/**
	 * 路由组请求
	 */
	userParty := app.Party("/users", func(context context.Context) {
		// 处理下一级请求
		context.Next()
	})
	/**
	 * 路由组下面的下一级请求
	 * ../users/register
	 */
	userParty.Get("/register", func(context context.Context) {
		app.Logger().Info("用户注册")
		context.HTML("<h1>用户注册功能</h1>")
	})

	/**
	 * 路由组下面的下一级请求
	 * ../users/login
	 */
	userParty.Get("/login", func(context context.Context) {
		app.Logger().Info("用户登录")
		context.HTML("<h1>用户登录功能</h1>")
	})

	userRouter := app.Party("/admin",userMiddleware)

	//Done
	userRouter.Done(func(context context.Context) {
		context.Application().Logger().Infof("response sent to " + context.Path())
	})

	userRouter.Get("/info", func(context context.Context) {
		context.HTML("<h1>用户信息</h1>")
		context.Next()// 手动显示调用
	})

	userRouter.Get("/query", func(context context.Context) {
		context.HTML("<h1>查询用户信息</h1>")
		//context.Next()
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
