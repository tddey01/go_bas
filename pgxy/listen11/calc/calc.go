package calc

import "fmt"

var (
	Sun int
)

func Add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

//导出变量量或函数。⾸首字⺟母⼤大写表示可导出，⼩小写表示私有。不不能被外部的包访问
//除了了可执⾏行行程序之外，⽤用户可以写⾃自定义包，⾃自定义包编译成静态库

func init() {
	fmt.Println("init func in calc 1")
}
