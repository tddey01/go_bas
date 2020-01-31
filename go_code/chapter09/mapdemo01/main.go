package main

import "fmt"

func main() {
	// map 的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" // ok？
	a["no2"] = "吴用" // ok？
	a["no1"] = "武松" //ok?
	a["no3"] = "吴用" //ok?
	fmt.Println(a)
	/*
				map 的基本介绍
				map 是 key-value 数据结构，又称为字段或者关联数组。类似其它编程语言的集合，在编程中是经常使用到
					var map 变量名 map[keytype]valuetype

				map 的声明
				 基本语法
				key 可以是什么类型
				golang 中的 map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还可以是只 包含前面几个类型的 接口, 结构体, 数组
				通常 key 为 int 、string
				注意: slice， map 还有 function 不可以，因为这几个没法用 == 来判断

				valuetype 可以是什么类型
				valuetype 的类型和 key 基本一样，这里我就不再赘述了
				通常为: 数字(整数,浮点数),string,map,struct

				map声明的举例:
				var a map[string]string
				var a map[string]int
				var a map[int]string
				var a map[string]map[string]string
				注意:声明是不会分配内存的，初始化需要 make ，分配内存后才能赋值和使用。


				对上面代码的说明
				1) map在使用前一定要make
				2) map的key是不能重复，如果重复了，则以最后这个key-value为准 3) map的value是可以相同的.
				4) map 的 key-value 是无序
				5) make内置函数数目

				func make
		func make(Type, size IntegerType) Type
		内建函数make分配并初始化一个类型为切片、映射、或通道的对象。其第一个实参为类型，而非值。make的返回类型与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：

		切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；
		     它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的切片。
		映射：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，这种情况下就会分配一个
		     小的起始大小。
		通道：通道的缓存根据指定的缓存容量初始化。若 size为零或被省略，该信道即为无缓存的。
	*/

	
}
