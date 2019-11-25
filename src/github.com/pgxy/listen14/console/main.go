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
}
