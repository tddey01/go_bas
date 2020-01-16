package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [5]int = [5]int{1, 5, 10, 6, 89}
	sort.Ints(a[:])
	fmt.Println("a", a) //证书排序

	var b [5]string = [5]string{"ab", "ec", "be", "fa", "ii"}
	sort.Strings(b[:])
	fmt.Println("b", b) //字符串排序

	var c [5]float64 = [5]float64{29.38, 22.32, 0.8, 99191.2}
	sort.Float64s(c[:])

	fmt.Println("c:", c) // 浮点数排序

}
