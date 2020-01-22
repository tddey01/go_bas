package main

import "fmt"

func main() {

	// 演示 for-range遍历数组
	var heroes [3]string = [3]string{"宋江", "吴用", "卢俊义"}
	fmt.Println(heroes)

	// 使用for循环遍历方式
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
	}
}
