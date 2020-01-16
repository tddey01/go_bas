package main

import (
	"strconv"
	"fmt"
)

func main() {
	// 统计字符串的长度 按字节 len(str)
	// golang 的编码统一为utf-8(ascii的字符(字母和数字)占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str  len", len(str)) //8

	str2 := "hello北京"
	// 字符串遍历， 同时处理有中文的问题，r:=[]rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数， n, err:= strconv.Atoi("12")
	n,err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误",err)
	}else {
		fmt.Println("转换结果是",n)
	}
}
