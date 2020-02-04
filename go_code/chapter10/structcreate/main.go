package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	//方式2 推荐使用第二种
	p2 := Person{"mary", 19}
	// p2.Name = "tom"
	// p2.Age = 18
	fmt.Println(p2)

	//方式3
	// 案例 var person *Person = new (Person)
	var p3 *Person = new(Person)
	//应为p3是一个指针  因此标准的给字段赋值方式
	// (*p3).Name = "smith"   也可以这样写 p3.Name = "smith"
	// 原因 go的设计者 为了程序员使用方便  在底层会对 p3.Name = "smith" 进行处理
	// 会给p3加上取值运算  	(*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "john"

	(*p3).Age = 30
	p3.Age = 100
	fmt.Println(*p3)

	// 第四种方式
	// 案例  var person *Person = &Person{}
	// var person *Person = &Person{"mary", 60} // 在括号里面可以直接赋值
	var person *Person = &Person{}
	// 因此person 是一个指针 ，因此标准的访问访问字段的方法
	// (*person).Name =  "scott"
	// go 的设置者为了程序员使用方便， 也可以 person.Name = "scott"
	// 原因和上面一样 底层会对 person.Name = "scott" 进行处理 会加上(*person)
	(*person).Name = "scott"
	person.Name = "scott~~"

	(*person).Age = 88
	person.Age = 10
	fmt.Println(*person)
}
