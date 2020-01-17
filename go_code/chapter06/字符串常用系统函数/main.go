package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度 按字节 len(str)
	//1 golang 的编码统一为utf-8(ascii的字符(字母和数字)占一个字节，汉字占用3个字节)
	str := "hello北"
	fmt.Println("str  len", len(str)) //8

	str2 := "hello北京"
	//2 字符串遍历， 同时处理有中文的问题，r:=[]rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//3 字符串转整数， n, err:= strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换结果是", n)
	}

	//4 整数转字符串 str = strvonv.Itoa(123456)
	str = strconv.Itoa(123456)
	fmt.Printf("str = %v str=%T\n", str, str)

	//5 字符串 转 []byte var  bytes = []byte("hello go")  类型
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//6 []byte 转成 字符串 str = string([]byte(97,98,99))
	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str)

	//7 十进制 二 八 十六 进制， str = strconb.FormatInt(123,2),返回对应的字符串
	str = strconv.FormatInt(123, 2) // 二进制
	fmt.Printf("123对应的二进制=%v str=%T\n", str, str)
	str = strconv.FormatInt(123, 8) //八进制
	fmt.Printf("123对应的十六进制=%v str=%T\n", str, str)
	str = strconv.FormatInt(123, 16) // 十六进制
	fmt.Printf("123对应的十六进制=%v str=%T\n", str, str)

	//8 查找子串是否指定字符串中 strings.Contains("seafood","foo") //true
	b := strings.Contains("seafood", "foo") //true
	fmt.Printf("b=%v b=%T\n", b, b)

	//9 统计一个字符串有几个指定的字符串， strings.Count("ceheese","e") //4
	num := strings.Count("ceheese", "h") //4
	fmt.Printf("num=%v num=%T\n", num, num)

	//10 不区分大小写字符串比较(==是区分字母大小的) fmt.Println(strings.EqualFold("abc","Abc")) // true
	s := strings.EqualFold("abc", "Abc") // strings.EqualFold 不区分大小写
	fmt.Println(s)

	fmt.Println("结果", "abc" == "Abc") // false  //区分字母大小写

	//11  返回字串在字符串第一次出现的index值， 如果没有返回-1 strings.Index("NLT_abc",abc) //4
	index := strings.Index("NLT_abc", "sss")
	fmt.Printf("NTL_abc 返回值=%v\n", index)

	//12 返回子串在字符串最后一次出现的index，如没有返回-1 : strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go") //3 字符串最后一次出现的位置
	fmt.Printf("idnex=%v\n", index)

	// 13 将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go语言", n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v  str2=%v\n", str, str2)

	// 14 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组： strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,wrold,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15)将字符串的字母进行大小写的转换: strings.ToLower("Go") // go strings.ToUpper("Go") // GO
	str = "goLang Hello"
	str = strings.ToLower(str) // 全部改变小写
	fmt.Println(str)

	str = strings.ToUpper(str) // 全部改变大写
	fmt.Println(str)

	//16 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn   ")
	str3 := " tn a lone gopher ntn    "
	str4 := strings.TrimSpace(str3) //字符串左右两边的空格去掉
	fmt.Printf("str4=%q\n", str4)

	//17)将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! ", " !")  // ["hello"] //将左右两边 ! 和 " "去掉
	str = strings.Trim("! hello! ", " !")
	fmt.Printf("str=%q\n", str)

	//18 将字符左边指定的字符去掉， strings.TrimLeft("! hello!", "!") //将左边！和""去掉
	str = strings.TrimLeft("! hello!", " !")
	fmt.Printf("str=%q\n", str)

	//19 将字符右边指定的字符去掉， strings.TrimLeft("! hello!", "!") //将右边！和""去掉
	str = strings.TrimRight("! hello ! ", "! ")
	fmt.Printf("str=%q\n", str)

	//20)判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v\n", b)

	//21)判断字符串是否以指定的字符串结尾: strings.HasSuffix("NLT_abc.jpg", ".jpg") // fales
	s = strings.HasSuffix("NLT_abc.jpg", "jpg")
	fmt.Printf("s=%v\n", s)
}
