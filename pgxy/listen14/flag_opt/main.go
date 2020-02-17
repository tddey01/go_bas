package main

import "flag"

import "fmt"

var recusive bool
var test string
var level int

func init() {

	flag.BoolVar(&recusive, "r", false, "recusive xxx")
	flag.StringVar(&test, "t", "default  string", "string option")
	flag.IntVar(&level, "l", 1, "level  of xxx")

	flag.Parse()

}

func main() {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("test %v\n", test)
	fmt.Printf("level %v\n", level)
}

/*
使⽤用flag包获取命令⾏行行参数
 ./flag_opt -r -t hello -l 24456
*/
