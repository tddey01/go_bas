package main

import "fmt"

// 递归
func test(n int) {
	if n > 2 {
		n-- // 死鬼
		test(n)
	} else {
		fmt.Println("n=", n)
	}
}
func main() {
	n := 4
	test(n)
}
