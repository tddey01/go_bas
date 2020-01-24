package main

import "fmt"

func main() {
	// 使用常规的for循环切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	// slice := arr[1:4] // 20,30,40
	// slice := arr[0:len(arr)]
	slice := arr[:] // 简写 两头都是默认值
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\n", i, slice[i])
	}

	// 使用for-range  方式遍历
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

}
