package main

import (
	"fmt"
	"strings"

	"github.com/pgxy/listen2/access"
)

func testBol() {
	var a bool
	fmt.Println(a)

	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)

	var b bool = true
	if a == true && b == true { //与
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	if a == true || b == true { // 或
		fmt.Println("|| right")
	} else {
		fmt.Println("||  not right")
	}
	fmt.Printf("%t %t\n", a, b) // 格式化输出

	//1. 布尔类型
	//a. var b bool 和 var b bool = true 和 var b = false
	//b. 操作符==和 != c. 取反操作符: !b
	//d. &&和||操作符
	//e. 格式化输出占位符: %t
}

func testInt() {
	var a int8 //取值范围在正128到-127之间
	a = 8
	fmt.Println("a=", a)
	a = -12
	fmt.Println("a=", a)

	var b int
	b = 138338338
	fmt.Println("b=", b)
	b = int(a)
	fmt.Println("b=", b)

	//var c float32 = 5.6
	var c = 5.6
	fmt.Println(c)

	fmt.Printf("a=%d a=%x c=%f\n", a, a, c)
	//2. 整数和浮点数类型
	//a. int8、int16、int32、int64
	//b. uint8、uint16、uint32、uint64
	//c. int 和 uint，和操作系统平台相关
	//d. float32 和 float64浮点类型
	//e. 所有整数 初始化为0，所有浮点数初始化为0.0，布尔类型初始化为false

	//3. 整数和浮点数类型
	//a. Go是强类型语⾔言，不不同类型相加以及赋值是不不允许的 b. 那怎么样才能实现，不不同类型相加呢?
	//c. 输出占位符:整数%d，%x⼗十六进制，%f浮点数
}

func testStr() {
	var a string = "htllo"
	fmt.Println(a)

	b := a // := 声明:=
	fmt.Println(b)

	c := "c:\nhello"
	fmt.Println(c)

	fmt.Printf("a=%v b=%v c=%v\n ", a, b, c)

	//5. 字符串串的两种表示⽅方式u
	//a. 双引号， “”，可以包含控制字符
	//b. 反引号， ``，所有字符都是原样输出
	c = `c:\nhello`
	fmt.Printf("a=%v b=%v c=%v\n ", a, b, c)

	//6. 字符串串常⽤用操作
	//字符串串
	//a. ⻓长度:len(str)
	//b. 拼接:+，fmt.Sprintf
	//c. 分割:strings.Split
	//d. 包含: strings.Contains
	//e. 前缀或后缀判断:strings.HasPrefix, strings.HasSuffix
	//f. ⼦子串串出现的位置: strings.Index(), strings.LastIndex()
	//g. join操作: strings.Join(a[]string, sep string)
	//var clen int
	//clen = len(c)
	clen := len(c)
	fmt.Printf("len of c = %d\n", clen)

	c = c + c //第一种拼接方式
	fmt.Printf("c = %s\n", c)

	c = fmt.Sprint("%s%s", c, c)
	fmt.Printf("c = %s\n", c)

	ips := "10.108.34.30;10.108.34.31"
	ipArray := strings.Split(ips, ";")
	fmt.Printf("first ip :%s\n", ipArray[0])
	fmt.Printf("second ip :%s\n", ipArray[1])

	result := strings.Contains(ips, "10.108.34.31") //判断一个串里面包含另一个字符串 是否为true
	fmt.Println(result)

	// 前缀或后缀判断:strings.HasPrefix, strings.HasSuffix
	str := "http://www.baidu.com"
	if strings.HasPrefix(str, "http") {
		fmt.Printf("str  is http url")
	} else {
		fmt.Printf("str is not http url")
	}

	if strings.HasSuffix(str, "com") {
		fmt.Printf("str  is com url")
	} else {
		fmt.Printf("str is not com url")
	}

	// ⼦子串串出现的位置: strings.Index(), strings.LastIndex()
	index := strings.Index(str, "baidu") // 从前往后
	fmt.Printf("baidu is idnex:%d\n", index)

	index = strings.LastIndex(str, "baidu")
	fmt.Printf("baidu last index:%d\n", index)

	// join操作: strings.Join(a[]string, sep string)
	var strArr []string = []string{"10.108.34.31", "10.108.34.32", "10.108.34.33"}
	result1 := strings.Join(strArr, ";")
	fmt.Printf("result=%s", result1)
}

func testOperator() {
	// 5. 操作符
	// a. 逻辑操作符， == 、 != 、<、<=、>=
	// b. 算数操作符， +、-、*、/、%
	var a int = 2
	if a == 2 {
		fmt.Printf("is right\n")
	} else {
		fmt.Printf("is  not right\n")
	}

	if a != 2 {
		fmt.Printf("is right\n")
	} else {
		fmt.Printf("is  not right\n")
	}
	a = a + 100
	fmt.Printf("a=%d\n", a)
}

func testAccess() {
	fmt.Printf("access.a= %d", access.A) // go 大写和小个 分为 公共 和私有类
}
func main() {
	// testBol()
	//testInt()
	// testStr()
	// testOperator()
	testAccess()

}
