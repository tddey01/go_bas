package main

import "fmt"

type MethodUtils struct {
	// 字段...
}

// 给MethodUtils 编写方法
func (mu MethodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 2、编写一个方法，提供m和n连个参数， 方法中答应一个m*n的矩形
func (mu MethodUtils) print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
// 3 、
编写一个方法算该矩形的面积(可以接收长len，和宽width)，
将其作为方法返回值。在main方法中调用该方法，接收返回的面积值并打印
*/

func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// 4 、 编写一个方法  判断一个数是奇数还是偶数
func (mu MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

func main() {
	/*
		1)编写结构体(MethodUtils)，编程一个方法，方法不需要参数，
		在方法中打印一个10*8 的矩形，在main方法中调用该方法。
	*/
	var mu MethodUtils
	mu.print()

	// 2、 练习
	fmt.Println()
	mu.print2(5, 20)

	areaRes := mu.area(2.5, 8.7)
	fmt.Println("面积等于 areaRes=", areaRes)

	// 4 、判断一个数是奇数还是偶数
	(&mu).JudgeNum(10)

}
