package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string `json:"monster_name"`
	Age      int    `json:"monster_age"`
	Birthday string `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string `json:"monster_skill"`
}

func testStruct() {
	// 对结构体进行序列化操作
	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Sal:      8000.00,
		Skill:    "牛魔拳",
	}
	//将monster数据序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

// 将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	// 使用map 需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "李家村"

	// 将a 这个map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后结果
	fmt.Printf("map 序列化后=%v\n", string(data))
}

// 对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用map前 需要先进行make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "17"
	//m2["address"] = "山东"
	m2["address"] = [2]string{"山东", "山西"}
	slice = append(slice, m2)

	// 将切片进行序列化
	// 将a 这个map进行序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后结果
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

// 对基本数据类型进行序列化
func testFloat64() {
	var num1 float64 = 2345.67
	//对num1 进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后结果
	fmt.Printf("num1 序列化后=%v\n", string(data))
}
func main() {
	// 演示结构体 map 切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
