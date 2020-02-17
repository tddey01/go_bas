package main

import "fmt"

type Ciscle struct {
	redius float64
}

// 2 声明一个犯法area和Circle绑定 可以放回面积
func (c *Ciscle) area() float64 {
	return 3.14 * c.redius * c.redius
}

// 为了提高效率  通常我们方法和结构体的指针类型绑定
func (c *Ciscle) area2() float64 {
	// 因此 c是指针， 因此我们标准的访问其自担的方式 (*c).redius
	// return 3.14 * (*c).redius * (*c).redius
	// (*c).redius 等价于 c.redius
	fmt.Printf("c 是 *Circle 指向地址=%p\n", c)
	c.redius = 10
	return 3.14 * c.redius * c.redius
}
func main() {
	// // 创建一个Circle结构体变量
	// var c Ciscle
	// c.redius = 4.0
	// res := c.area()
	// fmt.Println("res=", res)

	// 创建一个Circle 变量
	var c Ciscle
	fmt.Printf("main c 结构体变量 = %p\n", &c)
	c.redius = 7.0
	res2 := (&c).area2() // 标准访问方式
	// 编译器底层做了优化 (&c).area2() 等价于 c.area2()
	// 因为编译器会自动的给加上&c
	// fmt.Println("面积=", res2)

	res2 = c.area2()
	fmt.Println("面积=", res2)
	fmt.Println("c.radius=", c.redius)

}
