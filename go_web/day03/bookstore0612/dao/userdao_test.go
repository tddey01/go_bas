package dao

import (
	"fmt"
	"testing"
)

// 测试方法
func TestUser(t *testing.T) {
	fmt.Println("测试User")
	t.Run("验证用户名或密码:", testLogin)
	t.Run("验证用户名:", testRegist)
	//t.Run("验证保存用户:", testSavew)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取到用户信息", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息", user)
}
func testSavew(t *testing.T) {
	err := SaveUser("admin", "123456", "admin@localhost.com")
	if err != nil {
		fmt.Println("保存失败", err)
	}

}
