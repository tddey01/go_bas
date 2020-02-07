package main

import (
	"fmt"
	"go_bas/go_code/chapter11/encapexercise/model"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzsh1111", "999999", 40)
	if account != nil {
		fmt.Println("创建成功", account)
	} else {
		fmt.Println("创建失败")
	}
}
