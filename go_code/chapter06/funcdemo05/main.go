package main

import "fmt"

//在Go中，函数也是一种数据类型，
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型 因此在Go中，函数可以作为参数 并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 案例 2
//这时 myFun 就是 func(int, int) 类型
type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

// 编写一个sum函数
// 可变参数的使用
func sum(n1 int, args ...int) int {
	sum := n1
	//遍历args...
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0] 表示去除args 切片的第一个值  其他依次类推
	}
	return sum
}

func main() {

	a := getSum
	fmt.Printf("a的类型%T ，getSum类型%T\n", a, getSum)

	res := a(10, 40) // 等价于 res := getSum(10,40)
	fmt.Println("res=", res)

	//案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)

	// 给int 取别名 在go中 myInt和int虽然都是int类型  但是go认为myInt和int两个类型
	type myInt int

	var num1 myInt //
	var num2 int
	num1 = 40
	num2 = int(num1) //注意这里依然需要显示转换go认为myInt和int是两个类型
	fmt.Println("num1=", num1, "num2=", num2)

	// 案例2
	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//返回值命名 案例
	a1, b1 := getSumAndSub(1, 2)
	fmt.Printf("a1=%v b1=%v\n", a1, b1)

	// 测试 可变参数使用
	res4 := sum(10, 0, -1, 90, 10, 100)
	fmt.Println("res4=", res4)
}
