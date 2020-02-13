package main

import "fmt"

// Monkey结构体
type Monkey struct {
	Name string
}

//生命接口
type BirdAble interface {
	Flying()
}

func (this *Monkey) climbing() {
	fmt.Println(this.Name, "生来会爬树...")
}

type FishAble interface {
	Swimming()
}

// LittleMonkey 结构体
type LittleMonkey struct {
	Monkey // 匿名继承
}

// 让LittleMonkey实现BirdAble接口
func (this *LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习 会分享...")
}

func (this *LittleMonkey) Swimming() {
	fmt.Println(this.Name, "通过学习  会游泳...")
}

func main() {
	// 创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
