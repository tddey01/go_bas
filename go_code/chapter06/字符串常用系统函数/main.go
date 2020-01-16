package main

import (
	"fmt"
	"strconv"
	"strings"
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
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果是", n)
	}

	// 整数转字符串 str = strvonv.Itoa(123456)
	str = strconv.Itoa(123456)
	fmt.Printf("str = %v str=%T\n", str, str)

	// 字符串 转 []byte var  bytes = []byte("hello go")  类型
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	// []byte 转成 字符串 str = string([]byte(97,98,99))
	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str)

	// 十进制 二 八 十六 进制， str = strconb.FormatInt(123,2),返回对应的字符串
	str = strconv.FormatInt(123, 2) // 二进制
	fmt.Printf("123对应的二进制=%v str=%T\n", str, str)
	str = strconv.FormatInt(123, 8) //八进制
	fmt.Printf("123对应的十六进制=%v str=%T\n", str, str)
	str = strconv.FormatInt(123, 16) // 十六进制
	fmt.Printf("123对应的十六进制=%v str=%T\n", str, str)

	// 查找子串是否指定字符串中 strings.Contains("seafood","foo") //true
	b := strings.Contains("seafood", "foo") //true
	fmt.Printf("b=%v b=%T\n", b, b)

	// 统计一个字符串有几个指定的字符串， strings.Count("ceheese","e") //4
	num := strings.Count("ceheese", "h") //4
	fmt.Printf("num=%v num=%T\n", num, num)

	// 不区分大小写字符串比较(==是区分字母大小的) fmt.Println(strings.EqualFold("abc","Abc")) // true
	s := strings.EqualFold("abc", "Abc") // strings.EqualFold 不区分大小写
	fmt.Println(s)

	fmt.Println("结果", "abc" == "Abc") // false  //区分字母大小写

	// 返回字串在字符串第一次出现的index值， 如果没有返回-1 strings.Index("NLT_abc",abc) //4
	Index := strings.Index("NLT_abc", "sss")
	fmt.Printf("NTL_abc 返回值=%v\n", Index)
}
