package main

import "fmt"

// 将打印金字塔的代码封装到函数中
func printPyramid(totalLevel int) {
	// var totalLevel int = 10
	//i 表示层数
	for i := 1; i <= totalLevel; i++ {
		// 在打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		//j 表示每层答应多少*
		for j := 1; j <= 2*i-1; j++ {
			// if j == 1 || j == 2*i-1 || i == totalLevel {
			fmt.Print("*")
			// } else {
			// 	fmt.Print(" ")
			// }

		}
		fmt.Println()
	}
	// var num int = 9
	// for i := 1; i <= num; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d * %d = %d\t", j, i, j*i)
	// 	}
	// 	fmt.Println()
	// }
}
func main() {
	var n int
	fmt.Println("请输入打印金字塔的层数")
	fmt.Scanln(&n)

	printPyramid(n)
}
