package main

import "fmt"

func main() {
	// 假如 还有97天放假  问 xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d个星期零%d天\n", week, day)
	fmt.Println("helo")

}
