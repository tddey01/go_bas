package main

import "fmt"

// 编写一个函数 打印九九乘法表
func printMulti(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	var n int
	fmt.Println("请输入打印九九乘法表")
	fmt.Scanln(&n)

	printMulti(n)
}
