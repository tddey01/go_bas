package main

import "fmt"

func main() {
	// 第一种方式
	// map 的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" // ok？
	a["no2"] = "吴用" // ok？
	a["no1"] = "武松" //ok?
	a["no3"] = "吴用" //ok?
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	// 第三种方式
	// var heroes map[string]string = map[string]string{
	heroes := map[string]string{
		"hero1": "宋江1",
		"hero2": "卢俊义",
		"hero3": "李二牛",
	}
	heroes["hero4"] = "林冲"
	fmt.Println(heroes)
}
