package main

import "os"

import "fmt"

func main() {
	//os.Stdin //标准输入
	var buf [16]byte
	os.Stdin.Read(buf[:])
	// fmt.Println(string(buf[:]))
	os.Stdout.WriteString(string(buf[:]))
	fmt.Scanf()
	s = false
	/*
		终端其实是⼀一个⽂文件
		终端相关⽂文件的实例例
		os.Stdin:标准输⼊入的⽂文件实例例，类型为*File
		os.Stdout:标准输出的⽂文件实例例，类型为*File
		os.Stderr:标准错误输出的⽂文件实例例，类型为*File
	*/
}
