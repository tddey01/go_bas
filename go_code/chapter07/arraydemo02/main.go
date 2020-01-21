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
}
