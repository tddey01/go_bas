package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type Teacher struct {
	Name string `json:"name"`
	City string `json:"city"`
	Other string`json:"other"`
}

// 程序主入口
func main(){
	app := iris.New()
	app.Logger().SetLevel("debug")
	/**
	 * 通过Get请求
	 * 返回 WriteString
	 */
	app.Get("/getweb",func(context iris.Context){
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("hello world")
	})
	app.Get("/gethtml", func(context iris.Context) {
		context.HTML("html world")
	})

	//返回json数据
	type Student struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}
	app.Get("/getjson", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.JSON(iris.Map{"manages":"hello world","requestCode":200})
	})
	app.Get("/getStujson", func(context iris.Context) {
		context.JSON(Student{Name: "admin",Age: 20})
	})

	//返回xml数据
	type Peron struct {
		Name string `xml:"name"`
		Age int `xml:"age"`
	}
	app.Get("/getxml", func(context iris.Context) {
		context.XML(Peron{Name: "admin",Age: 29})
	})

	//返回Text
	app.Get("/gettext", func(context iris.Context) {
		context.Text("hello world" )
	})

	//PUT
	app.Put("/gethllo", func(context iris.Context) {
		path := context.Path()
		app.Logger().Info(path)
		context.WriteString("Put Request")
	})

	app.Put("/DeleteHello", func(context context.Context) {
		context.WriteString("PUT  Rquest DELETE")
	})


	// POST
	//app.Post("/user/userinfo", func(context context.Context) {
	//	//context.WriteString(" Post Request ")
	//	//user := context.Params().GetString("user")
	//	var tec Teacher
	//	err := context.ReadJSON(&tec)
	//	if err != nil{
	//		panic(err.Error())
	//	}else {
	//		app.Logger().Info(tec)
	//		context.WriteString(tec.Name)
	//	}
	//
	//})
	app.Post("/user/postinfo", func(context context.Context) {
		//context.WriteString(" Post Request ")
		//user := context.Params().GetString("user")

		var tec Teacher
		err := context.ReadJSON(&tec)
		if err != nil {
			panic(err.Error())
		} else {
			app.Logger().Info(tec)
			context.WriteString(tec.Name)
		}

		//fmt.Println(user)
		//teacher := Teacher{}
		//err := context.ReadForm(&teacher)
		//if err != nil {
		//	panic(err.Error())
		//} else {
		//	context.WriteString(teacher.Name)
		//}
	})

	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}