package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	//概念， 关于file的叫法
	// 1  file 对象
	// 2  file 指针
	// 3  file 文件句柄
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err= ", err)
	}
	//输入下文件 看看文件是什么  看出file 就是一个指针 *file
	fmt.Printf("file=%v\n", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
