package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	// 声明了两个没有实现的方法
	Start()
	Stop()
}
type Phone struct {
	name string
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
	name string
}

// 让iphone 实现 Usb 接口的方法
func (c Camer) Start() {
	fmt.Println("相机开始工作")
}
func (c Camer) Stop() {
	fmt.Println("相机停止工作")
}

func main(){
	// 定义一个Usb接口数组，可以存放Phone和Camear的结构变量
	// 这里就是体现出多态数组
	var UsbArr [3]Usb
	UsbArr[0] = Phone{"Vivo"}
	UsbArr[1] = Phone{"小米"}
	UsbArr[2] = Camer{"索尼"}

	fmt.Println(UsbArr)

}