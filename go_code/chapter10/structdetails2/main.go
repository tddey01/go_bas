package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"` // `json:"name"`  就是struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A 
	var b B
	// a = b    // 不能b的结构体赋值给a
	a = A(b) // 两个结构体 必须有相同的字段 名称 字段类型 字段个数 可以转换
	// ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	fmt.Println(a, b)

	// 1、 创建一个monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}package main

	import (
		"encoding/json"
		"fmt"
	)
	
	type A struct {
		Num int
	}
	
	type B struct {
		Num int
	}
	
	type Monster struct {
		Name  string `json:"name"` // `json:"name"`  就是struct tag
		Age   int    `json:"age"`
		Skill string `json:"skill"`
	}
	
	func main() {
		var a A 
		var b B
		// a = b    // 不能b的结构体赋值给a
		a = A(b) // 两个结构体 必须有相同的字段 名称 字段类型 字段个数 可以转换
		// ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
		fmt.Println(a, b)
	
		// 1、 创建一个monster变量
		monster := Monster{"牛魔王", 500, "芭蕉扇"}
	
		// 2 、 将monster变量序列化为 json 格式字符串
		// json.Marshal 函数中使用反射 者个讲解反射是详细介绍
		jsonStr, err := json.Marshal(monster)
		if err != nil {
			fmt.Println("json 错误处理", err)
		}
		fmt.Println("jsonSter", string(jsonStr))
	
	}
	

	// 2 、 将monster变量序列化为 json 格式字符串
	// json.Marshal 函数中使用反射 者个讲解反射是详细介绍
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 错误处理", err)
	}
	fmt.Println("jsonSter", string(jsonStr))

}
