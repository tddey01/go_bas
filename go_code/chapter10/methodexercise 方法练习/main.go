package main

import "fmt"

type Ciscle struct {
	redius float64
}

// 2 声明一个犯法area和Circle绑定 可以放回面积
func (c Ciscle) area() float64 {
	return 3.14 * c.redius * c.redius
}
func main() {
	// 创建一个Circle结构体变量
	var c Ciscle
	c.redius = 4.0
	res := c.area()
	fmt.Println("res=", res)
}
