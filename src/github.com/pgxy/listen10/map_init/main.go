package main

import "fmt"

func map_int1() {
	var a map[string]int = map[string]int{
		"stud01": 1000,
		"stud02": 2000,
		"stud03": 3000,
	}
	// 声明时进⾏行行初始化
	fmt.Println(a)
	a["stud01"] = 8888
	a["stud04"] = 3333
	fmt.Println(a)

	var key string = "stud04"
	fmt.Printf("the value of  key[%s] is :%d\n", key, a[key])
	//通过key访问map中的元素
}
func map_int2() {
	a := map[string]int{
		"stud01": 1000,
		"stud02": 2000,
		"stud03": 3000,
	}
	fmt.Println(a)
	a["stud04"] = 3394
	fmt.Println(a)

	b := "stud02"
	fmt.Printf("the value of  key[%s] is :%d\n", b, a[b])
}
func main() {
	map_int1() //第一种初始化方法
	map_int2() //第二种初始化方法
}
