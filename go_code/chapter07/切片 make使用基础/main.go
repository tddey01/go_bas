package main

import "fmt"

func main() {
	// 切片使用 make
	// var slice []float64 // 数组可以这样使用
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	// 对于切片 必须 make 使用
	fmt.Println(slice)
	fmt.Println("slice 的size=", len(slice))
	fmt.Println("sliec 的cap=", cap(slice))

	// 方式3
	// 第三种方式 定义一个切片 直接就指定具体数组  使用原理类似make的方式
	var strSlice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice)) //3
	fmt.Println("strSlice cap=", cap(strSlice))

}
