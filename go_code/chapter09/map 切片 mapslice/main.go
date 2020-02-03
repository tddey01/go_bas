package main

import "fmt"

func main() {
	// map 切片的使用
	/*
		要求 使用一个map来记录monster的信息 name和age， 也就是说一个monster对应一个map，并且妖怪的个数可以动态的增加=>map切片
	*/

	// 1、 声明一个map切片
	var monsters []map[string]string        // 切片
	monsters = make([]map[string]string, 2) // 准备放入两个妖怪
	// 2 、增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"

	}
	//
	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "400"

	}

	// 这里我们需要使用切片的append的函数，  可以动态的增加monster
	// 第一个不 先顶一个monster信息
	newMonster := map[string]string{
		"name": "行的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
