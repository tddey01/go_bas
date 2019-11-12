package main

import "fmt"

func main() {
	var a int32
	fmt.Printf("the addr of a:%p\n", &a)
	//变量量和内存地址
	a = 1000
	fmt.Printf("the addr of a:%p a:%d", &a, a)
	//每个变量量都有内存地址，可以说通过变量量来操作对应⼤大⼩小的内存
	//通过&符号可以获取变量量的地址
	//普通变量量存储的是对应类型的值，这些类型就叫值类型
	fmt.Println(a)
}
