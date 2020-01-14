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
func main() {

	a := getSum
	fmt.Printf("a的类型%T ，getSum类型%T\n", a, getSum)

	res := a(10, 40) // 等价于 res := getSum(10,40)
	fmt.Println("res=", res)

	//案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)

}
