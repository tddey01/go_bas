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

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func testClosure2() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))

}
func main() {
	//testClosure1()
	testClosure2()
}
