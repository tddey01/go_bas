package main

import (
	"fmt"
	"github.com/pgxy/listen6/calc"
)

// 可⻅见性，包内任何变量量或函数都是能访问的。包外的话，⾸首字⺟母⼤大写是可导出的 能够被其他包访问或调⽤用。⼩小写表示是私有的，不不能被外部的包访问
func main() {
	var s1 int = 100
	var s2 int = 200
	sum := calc.Add(s1, s2)
	fmt.Printf("s1+s2=%d\n", sum)

	fmt.Printf("calc A=%d\n", calc.A)
	fmt.Printf("calc a=%d\n", calc.Test())
}
