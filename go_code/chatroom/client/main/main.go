package main

import (
	"fmt"
	"go_bas/go_code/chatroom/client/process"
	"os"
)

//  定义两个变量 一个表示用户id 一个表示用户密码
var (
	userId   int
	userPwd  string
	userName string
)

func main() {
	// 接收用户的选择
	var key int
	// 判断是否还继续显示菜单
	// var loop = true
	for true {
		fmt.Println("------------------------欢迎登陆多人聊天室-------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 3 请选择(1-3)")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			// 1、 创建一个Userprocess的实例
			// loop = false
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			//2. 调用UserProcess，完成注册的请求、
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)

			// loop = false
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误 请重新输入")
		}
	}
	// 根据用户输入 显示新的提示信息
	// if key == 1 {
	// 	// 说明用户要登录
	// 	// fmt.Println("请输入用户id")
	// 	// fmt.Scanf("%d\n", &userId)
	// 	// fmt.Println("请输入用户密码")
	// 	// fmt.Scanf("%s\n", &userPwd)
	// 	// 把登录的函数 写到另一个文件内 比如写在login文件
	// 	// 这里我们需要重新调用
	// 	// login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登录失败")
	// 	// } else {
	// 	// 	fmt.Println("登录成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册逻辑")
	// }
}
