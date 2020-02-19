package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

// 将json字符串 反序列化成struct 结构体
func unmarshlStruct() {
	// 说明str  在项目开发中，是通过网络传输获取到  或者读取文件获取到
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"monster_birthday\":\"2011-11-11\",\"monster_sal\":8000,\"monster_skill\":\"牛魔拳\"}" // 先要转义

	// 定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	fmt.Printf("反序列化后 Monster=%v\n", monster)
}

//将map进行序列化
func testMap() string {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿~~~~~~"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将a这个map进行序列化
	//将monster 序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	//fmt.Printf("a map 序列化后=%v\n", string(data))
	return string(data)

}

//演示将json字符串，反序列化成map
func unmarshalMap() {
	//str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	str := testMap()
	//定义一个map
	var a map[string]interface{}

	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 map=%v\n", a)

}

// 将jons字符串反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"山东\",\"山西\"],\"age\":\"17\",\"name\":\"tom\"}]"

	//定义一个slice
	var slice []map[string]interface{}
	// 反序列化 反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshlStruct()
	unmarshalMap()
	unmarshalSlice()

}
