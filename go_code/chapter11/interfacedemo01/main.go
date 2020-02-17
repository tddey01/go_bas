package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
}

// 让iphone 实现 Usb 接口的方法
func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}

//相机实现 Usb接口方法

type Camer struct {
}

// 让iphone 实现 Usb 接口的方法
func (c Camer) Start() {
	fmt.Println("相机开始工作")
}
func (c Camer) Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type Computer struct {
}

//编写一个方法Working 方法 接收Usb接口类型变量
// 只要是实现了 Usb接口 所谓实现Usb接口 就是指实现了 Usb 接口声明所有方法
func (c Computer) Working(usb Usb) {
	// 通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	// 测试
	// 先创建结构体
	computer := Computer{}
	phone := Phone{}
	camer := Camer{}

	// 关键点
	computer.Working(phone)
	computer.Working(camer)
}

// 接口就是引用类型
