package main

import "fmt"

func testPoint1() {
	var a int32
	a = 1000
	fmt.Printf("the addr of a:%p a:%d\n", &a, a)

	var b *int32 // 指针类型变量量的默认值为nil，也就是空地址
	fmt.Printf("the  addr b :%p b:%v\n", &b, b)

	b = &a
	fmt.Printf("the  addr b :%p b:%v\n", &b, b)
	//通过* 符号可以获取指针变量量指向的变量量
}

func testPoint2() {
	var a int = 200
	var b *int = &a

	fmt.Printf("b指向的地址存储的值为:%d\n", *b)

	*b = 1000
	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	fmt.Printf("b指向的地址存储的值为:%d\n", a)
}

func main() {
	//testPoint1()
	testPoint2()
}
