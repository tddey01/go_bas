package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom"      // ok
	fs := reflect.ValueOf(&str) // ok fs --> string
	// fs.SetString("jack")        //error
	fs.Elem().SetString("jack") //必须这种改值
	// reflect 修改必须 传入指针  使用fs.Elem().SetString("jack") 开头的fs.Elem()函数 才可以修改值
	fmt.Printf("%v\n", str)
}
