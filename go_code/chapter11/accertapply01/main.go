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
	fmt.Println("手机开始工作...")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作...")
}

func (p Phone) Call() {
	fmt.Println("手机 打电话...")
}

//相机实现 Usb接口方法
type Camer struct {
	name string
}

// 让iphone 实现 Usb 接口的方法
func (c Camer) Start() {
	fmt.Println("相机开始工作...")
}
func (c Camer) Stop() {
	fmt.Println("相机停止工作...")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言..[注意体会!!!]
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	// 定义一个Usb接口数组，可以存放Phone和Camear的结构变量
	// 这里就是体现出多态数组
	var UsbArr [3]Usb
	UsbArr[0] = Phone{"Vivo"}
	UsbArr[1] = Phone{"小米"}
	UsbArr[2] = Camer{"索尼"}

	//遍历usbArr
	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb 接口声明的方法外，还需要调用Phone 特有方法 call. =》类型断言
	var computer Computer
	for _, v := range UsbArr {
		computer.Working(v)
		fmt.Println()
	}

	//fmt.Println(UsbArr)

}
