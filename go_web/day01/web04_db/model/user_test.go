package model

import (
	"fmt"
	"testing"
)

// TestMain =函数可以在测试函数执行之前做一些其他的操作
func testMain(m *testing.M) {
	fmt.Println("测试开始")
	//通过m.Run()来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试 User 中的相关方法")

	// 测试子函数  通过 t.Run() 来之执行测试函数
	// t.Run ("测试添加用户:", testAddUser)
	//t.Run("测试查询获取用户", testUser_GetUserById)
	t.Run("测试获取User所有用户", testGetUserAll)
}

// 如果函数名不是以Test开头， 那么函数默认不执行， 我们可以将它设置为一个自测试函数
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户测试子函数")
	// fmt.Println("测试添加用户")
	// user := &User{}
	// // 调用添加用户的方法
	// user.AddUser()
	// user.AddUser2()
}

// 测试函数
func testUser_GetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := User{
		Id: 3,
	}
	//	 调用获取User的方法
	u, _ := user.GetUserById()
	fmt.Println("得到的User的信息", u)
}

// 测试 所有用户 User
func testGetUserAll(t *testing.T) {
	fmt.Println("查询所有记录")
	user := &User{}
	//	 调用获取所有User方法
	us, _ := user.GetUserAll()
	//	遍历切片
	for k, v := range us {
		fmt.Printf("第%v个用户%v\n", k+1, v)
	}
}
