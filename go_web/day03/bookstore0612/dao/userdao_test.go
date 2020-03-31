package dao

import (
	"fmt"
	"go_bas/go_web/day03/bookstore0612/model"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试book中的方法")
	m.Run()
}

// 测试方法
func TestUser(t *testing.T) {
	// fmt.Println("测试User")
	// t.Run("验证用户名或密码:", testLogin)
	// t.Run("验证用户名:", testRegist)
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

func TestBook(t *testing.T) {
	fmt.Println("Book 测试bookdao中的相关函数")
	// t.Run("测试获取所有的图书", testGetBook)
	//t.Run("测试获取所有的图书", testAddBook)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取图书", testGetBooks)
	t.Run("测试更新图书", TestUpdateBooks)
}

func testGetBook(t *testing.T) {
	books, _ := GetBook()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, *v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.8,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(book)
}

func testDeleteBook(t *testing.T) {
	//调用删除图书的函数
	DeleteBook("35")
}

func testGetBooks(t *testing.T) {
	//调用获取图书的函数
	book, _ := GetBookByID("31")
	fmt.Println(book)
}

func TestUpdateBooks(t *testing.T) {
	book := &model.Book{
		Id:      31,
		Title:   "西游记",
		Author:  "罗贯中",
		Price:   110.02,
		Sales:   90,
		Stock:   200,
		ImgPath: "static/img/default.jpg",
	}
	// 调用更新图书的函数
	UpdateBook(book)
}
