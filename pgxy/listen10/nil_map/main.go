package main

import "fmt"

// 初始化map

func main() {
	var a map[string]int
	fmt.Printf("a:%v\n", a)
	//map类型是⼀一个key-value的数据结构。
	// 声明和初始化

	// a["stud01"] = 100 声明玩 不能直接操作 必须初始化
	//map必须初始化才能使⽤用，否则panic

	if a == nil { //map类型的变量量默认初始化为nil，需要使⽤用make分配map内存
		a = make(map[string]int, 16)
		fmt.Printf("a %v\n", a)
		a["stu01"] = 1000
		a["stu02"] = 1000
		a["stu03"] = 1000
		//map插⼊入操作
		fmt.Printf("a=%v\n", a)
	}

}
