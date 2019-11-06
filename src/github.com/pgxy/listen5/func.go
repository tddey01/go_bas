package main

import "fmt"

func sayhello() {
	fmt.Printf("hellow word\n")
}

func add(a, b int) int {
	// 连续的⼀一系列列参数的类型是⼀一样，前⾯面的类型可以不不写
	sum := a + b
	return sum
}

func main() {
	//sayhello()
	s := add(100, 300)
	fmt.Println(s)
}
