package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b string
	var c float32
	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	//fmt.Fscan(os.Stdin, &a, &b, &c)
	//fmt.Scanf("%d%s%f", &a, &b, &c)
	//fmt.Println(a, b, c)
	fmt.Fprintln(os.Stdout, "stdout", a, b, c)

	/*
		终端其实是⼀一个⽂文件
		终端相关⽂文件的实例例
		os.Stdin:标准输⼊入的⽂文件实例例，类型为*File
		os.Stdout:标准输出的⽂文件实例例，类型为*File
		os.Stderr:标准错误输出的⽂文件实例例，类型为*File
	*/
}
