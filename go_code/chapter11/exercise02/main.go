package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3 实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 1. 按年龄的从小到大排序
func (hs HeroSlice) Less(i, j int) bool { // 布尔类型判断
	//return hs[i].Age > hs[j].Age   // 降序排序
	//return hs[i].Age < hs[j].Age   // 升序排序
	// 改成对Name排序
	return hs[i].Name < hs[j].Name // 升序排序
}

func (hs HeroSlice) Swap(i, j int) {
	// //交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 下面一句话等价于上面三居话
	hs[i], hs[j] = hs[j], hs[i]
}

//1.声明Student结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//将Student的切片，安Score从大到小排序!!

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	// 要求对 intlice切片进行排序
	// 1 冒泡排序
	// 2 可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//请大家对结构体切片进行排序
	//1. 冒泡排序...
	//2. 也可以使用系统提供的方法

	// 测试看看我们是否可以对接口结构体切片排序
	var herose HeroSlice // 切片类型
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 将 hero append 到heroes切片
		herose = append(herose, hero)
	}
	// 看看排序前的顺序
	for _, v := range herose {
		fmt.Println(v)
	}

	// 调用sort.Sort
	sort.Sort(herose)
	fmt.Println("--------------------------------")
	// 看看排序后的顺序
	for _, v := range herose {
		fmt.Println(v)
	}

	i := 10
	j := 20
	i, j = j, i
	fmt.Println("i=", i, "j=", j)
}
