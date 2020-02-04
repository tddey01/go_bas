package main

import "fmt"

// 如果结构体的字段类型： 指针、slice、和map的零值都是nil，即还没有分配空间
// 如果需要使用这样的字段， 需要先make，才能使用

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针
	slice  []int             // 切片
	map1   map[string]string // 切片
}

func main() {
	// 定义结构体变量
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.slice == nil {
		fmt.Println("ok2")
	}

	if p1.map1 == nil {
		fmt.Println("ok3")
	}

	// 使用slice  再次声明  使用前 一定要make
	p1.slice = make([]int, 10)
	p1.slice[0] = 100 // ok？ 报错

	// 同样使用map 一定要先使用make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	p1.Age = 10

	p1.ptr = &p1.Age //指针

	fmt.Println(p1)

}
