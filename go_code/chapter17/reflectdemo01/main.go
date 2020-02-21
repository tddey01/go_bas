package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest01(b interface{}) {
	// 通过反射获取传入的变量的 type 、kind 值
	// 1、 先获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("reflct=,", rType)

	// 2 获取到reflect。value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	// 下面我们将 rval 转成interface{} 类型
	iv := rVal.Interface()

	// 将interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println(num2)

}

// 专门演示反射【对结构体的反射】
func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rval := reflect.ValueOf(b)

	//3、  获取 变量对应的Kind
	// 1、rVal.kind() -->
	kind1 := rval.Kind()
	// 2、rType.Kind() -->
	kind2 := rType.Kind()
	fmt.Printf("Kind1 =%v kind2 =%v\n", kind1, kind2)

	iv := rval.Interface()
	fmt.Printf("iv=%v iv=%T\n", iv, iv)
	// 将interface{} 通过断言转成需要的类型

	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	// 编写一个案例
	// 对(基本数据类型、interface{}、reflect.value) 进行反射基本操作

	// 1、 先定义一个int
	// var num int = 100
	// reflectTest01(num)

	// 定义一个student实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)

}
