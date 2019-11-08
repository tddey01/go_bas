package main

import "fmt"

func testMake1() {
	var a []int //[]中没有⻓长度

	a = make([]int, 5, 10)
	// 第一个参数5 代表长度是5
	// 第二个参数10  代表是容量

	a[0] = 10
	//a[1] = 20

	//a = append(a,11)
	//// append(a，11) 扩容
	//fmt.Printf("a=%v\n", a)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	//观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

func testMake2() {
	var a []int

	a = make([]int, 5, 10)
	// a[5] = 100  // 不能越界访问
	a = append(a, 10)
	fmt.Printf("a=%v\n", a)

	b := make([]int, 0, 10)
	fmt.Printf("a=%v len:%d  cap:%d\n", b, len(b), cap(b))
	b = append(b, 100)
	fmt.Printf("a=%v len:%d  cap:%d\n", b, len(b), cap(b))
}

func testCap() {
	a := [...]string{"a", "b", "c", "d", "d", "f", "g", "h"}
	b := a[1:3]
	fmt.Printf("b:%v len:%d cap:%d \n", b, len(b), cap(b))

}
func main() {
	// testMake1()
	// testMake2()
	testCap()
}
