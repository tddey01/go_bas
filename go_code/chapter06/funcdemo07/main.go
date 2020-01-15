package main

import "fmt"

// 累加器
func AddUpeer() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=", str) // 1 str="hello" 2
		return n
	}
}

func main() {
	//使用前面的代码
	f := AddUpeer()
	fmt.Println(f(0))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
