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

func calc(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
	// 对返回值进⾏行行命名
}
func calc_v1(b ...int) int {
	// 接收0个或多个int的参数
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func calc_v2(a int, b ...int) int {
	// 接收0个或多个int的参数
	sum := a
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func main() {
	//sayhello()
	//s := add(100, 300)
	//fmt.Println(s)

	//sum, _ := calc(100, 200)
	//fmt.Println(sum)

	//sum := calc_v1()
	//sum := calc_v1(10)
	//sum := calc_v1(10,20,30,40,50)
	//fmt.Printf("sum-%d\n",sum)

	sum := calc_v2(10)
	fmt.Printf("sum-%d\n", sum)
}
