package main

import (
	"fmt"
)

func testSlice0() {
	var a []int //定义⼀一个int类型的空切⽚片
	if a == nil {
		fmt.Printf("a is nil\n")
	} else {
		fmt.Printf("a = %v\n", a)
	}
	a[0] = 100
}

func testSlice1() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4] //基于数组a 创建⼀一个切⽚片，包括元素a[1] a[2] a[3]
	fmt.Printf("slice b:%v\n", b)
	fmt.Printf("b[0]=%d\n", b[0])
	fmt.Printf("b[1]=%d\n", b[1])
	fmt.Printf("b[2]=%d\n", b[2])
	fmt.Printf("b[3]=%d\n", b[3])
}

func testSlice2() {
	a := []int{1, 2, 3, 4, 5} // 创建⼀一个数组并返回⼀一个切⽚片

	fmt.Printf("slice a:%v type of a:%T\n", a, a)
}

func testSlice3() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Printf("slice b:%v\n", b)

	// c := a[1:len(a)]
	c := a[1:]
	fmt.Printf("slice c:%v\n", c)
	//d := a[0:3]
	d := a[:3]
	fmt.Printf("slice d:%v\n", d)
	// e  := a[0:len(a)]
	e := a[:]
	fmt.Printf("slice e:%v\n", e)

	//a) arr[start:end]:包括start到end-1(包括end-1)之间的所有元素
	//b) arr[start:]:包括start到arr最后⼀一个元素(包括最后⼀一个元素)之间的所有元素
	//c) arr[:end]:包括0到end-1(包括end-1)之间的所有元素
	//d) arr[:]:包括整个数组的所有元素
}

func testSlice4() {
	a := [...]int{1, 2, 3, 4, 5, 7, 8, 9, 11} // 创建⼀一个数组，其中[...]是编译器器确定数组的⻓长度,darr的⻓长度是

	fmt.Printf("array a:%v type of a:%T\n", a, a)
	b := a[2:5] //基于darr创建⼀一个切⽚片dslice,包括darr[2],darr[3],darr[4]三个元素
	fmt.Printf("slice b:%v type of b:%T\n", b, b)
	/*
		b[0] = b[0] + 10
		b[1] = b[1] + 20
		b[2] = b[2] + 30
	*/
	/*
		for index, val := range b {
			fmt.Printf("b[%d]=%d\n", index, val)
		}
	*/
	for index := range b {
		//对于dslice中每个元素进⾏行行+1，其实修改是darr[2],darr[3],darr[4]
		b[index] = b[index] + 10
	}
	fmt.Printf("after modify slice b, array a:%v type of a:%T\n", a, a)
}

func testSlice5() {
	//切⽚片修改
	a := [...]int{1, 2, 3}
	//创建⼀一个切⽚片，包含整个数组的所有元素
	s1 := a[:]
	s2 := a[:]
	s1[0] = 100
	fmt.Printf("a=%v s2=%v\n", a, s2)
	s2[1] = 200
	fmt.Printf("a=%v s1=%v\n", a, s1)
}

func main() {
	// testSlice0()
	//testSlice1()
	//testSlice2()
	// testSlice3()
	testSlice4()
	//testSlice5()
}
