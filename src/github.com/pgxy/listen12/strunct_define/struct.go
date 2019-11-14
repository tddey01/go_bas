package main

import "fmt"

type User struct {
	// 对象属性
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user User
	user.Age = 10
	user.AvatarUrl = "https://www.google.com"
	user.Sex = "男"
	user.Username = "user01"

	fmt.Printf("user.usernmae=%s  age=%s  sex=%s  avatar=%s\n", user.Username, user.Age, user.Sex, user.AvatarUrl)

	var user2 User = User{
		//struct初始化⽅方法2
		Username: "user02",
		// Age:       18,
		Sex: "女",
		// AvatarUrl: "https://www.baidu.com",
	}
	fmt.Printf("user2 = %#v\n", user2)
	// 更简洁方法方法3
	user3 := User{
		Username: "user03",
		// Age:       18,
		Sex:       "女",
		AvatarUrl: "https://www.baidu.com/a.jpg",
	}
	fmt.Printf("user3 = %#v\n", user3)
}

//⾯面向对象编程1
/*
struct声明和定义
struct初始化⽅方法1
*/
