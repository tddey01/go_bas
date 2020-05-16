package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

func main() {
	// 1  创建数据库对象
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/cms?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	// 2 数据库引擎关闭
	defer engine.Close()


	// 数据库引擎查询
	engine.ShowSQL(true)                     // 设置显示SQL
	engine.Logger().SetLevel(core.LOG_DEBUG) // 设置日志显示
	engine.SetMaxIdleConns(10)               //设置最大连接数

	// 查询表的所有数据
	session := engine.Table("user")
	count, err := session.Count()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(count)

	//使用原生sql语句进行查询
	result, err := engine.Query("select  * from user")
	if err != nil {
		panic(err.Error())
	}
	for key, value := range result {
		fmt.Println(key, value)
	}
}

type Person struct {
	Age int
	Name string
}

func OrmMapping() {

}
