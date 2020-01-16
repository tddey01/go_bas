package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello world\n"
	err := ioutil.WriteFile("./test.dat", []byte(str), 0755)
	if err != nil {
		fmt.Println("write file failed ,err :", err)
		return
	}
}
