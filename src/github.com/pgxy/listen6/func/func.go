package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func testFunc1() {
	f1 := add
	fmt.Printf("type of f1=%T\n", f1) // 打印变量的类型
	sum := f1(2, 5)
	fmt.Printf("sum=%d\n", sum)
}

func testFunc2() {
	f1 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("type of %T\n", f1)
	sum := f1(5, 7)
	fmt.Printf("sum=%d\n", sum)
}

func main() {
	testFunc1()
	testFunc2()
}
