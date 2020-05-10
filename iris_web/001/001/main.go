package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	// 创建app结构体对象
	//app := iris.New()
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris! vx1"})
	})
	// 处理get请求，请求的url为：/getRequest
	app.Get("/getRequest", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
	})
	//1、处理Get请求
	app.Get("/userpath", func(context context.Context) {
		//获取Path
		path := context.Path()
		app.Logger().Info(path)
		//写入返回数据：string类型
		context.WriteString("请求路径" + path)

	})

	// 2 .处理Get请求 并接受参数
	app.Get("/userinfo", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//获取get请求所携带的参数
		userName := context.URLParam("username")
		app.Logger().Info(userName)
		Pwd := context.URLParam("pwd")
		app.Logger().Info(Pwd)

		//返回html数据格式
		//context.HTML("<h1>" + userName + "," + Pwd + "</h1>")
		context.HTML("<h1>" + userName + "," + Pwd + "</h1>")
	})

	// 3 处理Post请求 form表单的字段获取
	app.Post("/postlogin", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)
		//context.PostValue方法来获取post请求所提交的for表单数据
		name := context.PostValue("name")
		pwd := context.PostValue("pwd")
		app.Logger().Info(name, "  ", pwd)
		context.HTML(name)
	})
	//4、处理Post请求 Json格式数据
	/**
	 * Postman工具选择[{"key":"Content-Type","value":"application/json","description":""}]
	 * 请求内容：{"name": "davie","age": 28}
	 */
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	app.Post("/postjson", func(context context.Context) {
		// 1 path
		path := context.Path()
		app.Logger().Info("请求URL：", path)
		// 2 解析JOSN
		var person Person
		// context.ReadJSON()
		if err := context.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		//输出：Received: main.Person{Name:"davie", Age:28}
		context.Writef("ReadJSON:%#+v\n", person)
		app.Logger().Info(person)

	})
	//5.处理Post请求 Xml格式数据
	/**
	 * 请求配置：Content-Type到application/xml（可选但最好设置）
	 * 请求内容：
	 *
	 *  <student>
	 *		<stu_name>davie</stu_name>
	 *		<stu_age>28</stu_age>
	 *	</student>
	 *
	 */
	type Student struct {
		//XMLName xml.Name `xml:"student"`
		StuName string `xml:"name"`
		StuAge  int    `xml:"age"`
	}
	app.Post("/postxml", func(context context.Context) {
		// 1 path
		path := context.Path()
		app.Logger().Info("请求URL: ", path)
		// 2 解析xml
		var student Student
		if err := context.ReadXML(&student); err != nil {
			panic(err.Error())
		}
		//输出：
		context.Writef("ReadXML: %#+v\n", student)
		app.Logger().Info(student)
	})
	app.Put("/userput", func(context context.Context) {
		path := context.Path()
		app.Logger().Info(path)

	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	// 端口监听
	//app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	////第一种方法
	//application.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	//// 第二种方法
	//application.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}
