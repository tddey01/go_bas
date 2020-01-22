package main

import "fmt"

func main() {
	var intArr [3]int // int 占用8个字节
	// 当我们定义完数组后，其实数组的各个元素有默认值 0
	fmt.Println(intArr)

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Printf("intarr的地址=%p intArr[0] 地址%p intArr[1] 地址%p \n", &intArr[0], &intArr[1], &intArr[2]) // 获取内存地址
	//intarr的地址=0xc0000141a0 intArr[0] 地址0xc0000141a8 intArr[1] 地址0xc0000141b0

	//从终端循环输入5个成绩， 保存到float64数组，并输出
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入%d个元素的值", i+1)
		fmt.Scanln(&score[i])
	}
	// 变量数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
}
