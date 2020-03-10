package model

import (
	"fmt"
	"testing"
)

func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户")
	user := &User{}
	// 调用添加用户的方法
	user.AddUser()
	user.AddUser2()
}
