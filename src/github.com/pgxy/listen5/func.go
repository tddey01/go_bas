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

func calc(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	//sayhello()
	//s := add(100, 300)
	//fmt.Println(s)

	sum, sub := calc(100, 200)
	fmt.Println(sub, sum)
}
