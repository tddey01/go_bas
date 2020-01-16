package main

import "fmt"

//一个函数 test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1) // 输出结果=？11
}

// 一个函数 getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ", sum)
	// 当函数有return语句时， 就是将结果返回给调用者
	// 即谁调用我， 就返回给谁
	return sum
}

// 编写要给函数，可以一计算两个数的和和差 并返回结果
func getSumAndsub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main() {
	//调用test
	n1 := 10
	test(n1)
	fmt.Println("n1=", n1)
	// n2 := 11
	sum := getSum(10, 20)
	fmt.Println("main sum = ", sum)

	// 调用getSumAndsub
	res1, res2 := getSumAndsub(1, 2) // res1 = 3 res2 = -1
	fmt.Println("res1=", res1, "res2=", res2)

	//希望忽略某个返回值，则使用_ 符号表示占位忽略
	_, res3 := getSumAndsub(3, 9)
	fmt.Println("res3=", res3)
}
