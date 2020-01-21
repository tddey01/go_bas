package main

import "fmt"

func main() {
	var intArr [3]int // int 占用8个字节
	// 当我们定义完数组后，其实数组的各个元素有默认值 0
	fmt.Println(intArr)

	fmt.Printf("intarr的地址=%p intArr[0] 地址%p intArr[1] 地址%p \n", &intArr, &intArr[0], &intArr[1]) // 获取内存地址
}
