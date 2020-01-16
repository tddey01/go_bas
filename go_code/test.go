package main

import "fmt"

var name = "tom" // 全局变量
func test01() {
	fmt.Println(name) //tom jack
}
func test02() { //编译器采用就近原则
	name := "jack"
	fmt.Println(name) // jack
}
func main() {
	fmt.Println(name) //tom
	test01()          // tom
	test02()          // jack
	test01()          // jack

}
