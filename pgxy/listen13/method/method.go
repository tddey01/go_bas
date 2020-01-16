package main

import "fmt"

type People struct { // 定义一个结构 一个对象
	Name    string
	Country string
}

func (p People) Print() { // 前面接收者  传参给p
	fmt.Printf("name =  %s  country = %s\n", p.Name, p.Country)
}

func (p People) Set(name string, country string) {
	p.Name = name
	p.Country = country
}

func (p *People) Setv2(name string, country string) { //指针函数
	p.Country = country
	p.Name = name
}

func main() {
	var p1 People = People{
		Name:    "people01",
		Country: "china",
	}
	p1.Print()
	p1.Set("peple02", "enligsh")
	p1.Print()

	//(&p1).Setv2("people02","enligsh") // 可以成功修改
	p1.Setv2("p2", "c2")
	p1.Print()
}
