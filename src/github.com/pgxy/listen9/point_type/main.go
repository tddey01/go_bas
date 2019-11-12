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

func modify(a *int) {
	*a = 100
	//指针变量量传参
}

func testPoint3() {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b:%d\n", b)
}

func modify_arr(a *[3]int) {
	(*a)[0] = 100 // 下标操作
}
func testPoint4() {
	var b [3]int = [3]int{1, 3, 2}
	modify_arr(&b)
	fmt.Printf("b:%v\n", b)
	//指针变量量传参示例例2
}

func testPoint5() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("a=:%d\n", *a)

	var b *[]int = new([]int)
	fmt.Printf("*b = %v\n", *b)
	// (*b)[0] =  100 错误的切片
	(*b) = make([]int, 5, 100)
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b :%v\n", *b)
}

func main() {
	//testPoint1()
	// testPoint2()
	// testPoint3() //指针变量量传参
	// testPoint4() //指针变量量传参示例例2
	testPoint5()
}
