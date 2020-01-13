package main

import "fmt"

//在Go中，函数也是一种数据类型，
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以
func getSum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {

	a := getSum
	fmt.Printf("a的类型%T ，getSum类型%T\n", a, getSum)

	res := a(10, 40) // 等价于 res := getSum(10,40)
	fmt.Println("res=", res)

}
