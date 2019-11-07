package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func testClosure1() {
	f := Adder()
	ret := f(1)
	fmt.Printf("f(1) ret=%d\n", ret)
	ret = f(20)
	fmt.Printf("f(20) ret=%d\n", ret)
	ret = f(100)
	fmt.Printf("f(100) ret=%d\n", ret)

	f1 := Adder() //重新调用闭包 代表重新开始
	ret = f1(1)
	fmt.Printf("f1(1) ret=%d\n", ret)
}

func main() {
	testClosure1()

}
