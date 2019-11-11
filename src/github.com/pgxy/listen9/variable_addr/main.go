package main

import "fmt"

func main() {
	var a int32
	fmt.Printf("the addr of a:%p\n", &a)
	//变量量和内存地址
	a = 1000
	fmt.Printf("the addr of a:%p a:%d", &a, a)
}
