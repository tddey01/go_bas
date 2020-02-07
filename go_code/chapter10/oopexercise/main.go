package main

import "fmt"

/*
学生案例：
编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

*/
type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 Name=[%v] Gender=[%v] Age=[%v] Id=[%v] Score=[%v]",
		student.Name, student.Gender, student.Age, student.Id, student.Score)
	return infoStr
}

/*
1)编程创建一个Box结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终端获取
2)声明一个方法获取立方体的体积。
3)创建一个Box结构体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

// 景区门票案例

// 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
// 请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("考虑到安全，就不要游玩")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Println("免费")
	}
}

func main() {
	//测试
	// 创建一个student实例变量
	var stu = Student{
		Name:   "tom",
		Gender: "male",
		Age:    22,
		Id:     1000,
		Score:  99.98,
	}
	str := stu.say()
	fmt.Println(str)

	//测试变量
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volumn := box.getVolumn()
	fmt.Printf("体积为=%.2f\n", volumn)

	//测试
	var v Visitor
	for {
		fmt.Println("请你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
