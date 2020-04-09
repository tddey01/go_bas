package dao

import (
	"fmt"
	"go_bas/go_web/day05/bookstore0612/model"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	// fmt.Println("测试userdao中的函数")
	// t.Run("验证用户名或密码：", testLogin)
	// t.Run("验证用户名：", testRegist)
	// t.Run("保存用户：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@atguigu.com")
}

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBook)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取一本图书", testGetBook)
	//t.Run("测试更新图书", testUpdateBook)
	//t.Run("测试获取带分页的图书", testGetPageBooks)
	//t.Run("测试获取带分页和价格范围的图书", testGetPageBooksByPrice)

}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%v本图书的信息是：%v\n", k+1, v)
	}
}
func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   88.88,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(book)
}
func testDeleteBook(t *testing.T) {
	//调用删除图书的函数
	DeleteBook("34")
}
func testGetBook(t *testing.T) {
	//调用获取图书的函数
	book, _ := GetBookByID("32")
	fmt.Println("获取的图书信息是：", book)
}
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      32,
		Title:   "3个女人与105个男人的故事",
		Author:  "罗贯中",
		Price:   66.66,
		Sales:   10000,
		Stock:   1,
		ImgPath: "/static/img/default.jpg",
	}
	//调用更新图书的函数
	UpdateBook(book)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书的信息是：", v)
	}
}
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("3", "10", "30")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数是：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	fmt.Println("当前页中的图书有：")
	for _, v := range page.Books {
		fmt.Println("图书的信息是：", v)
	}
}

func TestSession(t *testing.T) {
	fmt.Println("测试Sesseion相关方法")
	//t.Run("测试添加Session",testAddSession)
	t.Run("测试查看Session",testGetSessionById )
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "13838381438",
		UserName:  "马蓉",
		UserID:    1,
	}

	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("13838381438")
}

func testGetSessionById(t *testing.T) {
	sess,_ := GetSessionById("ee39f56d-0594-4e0d-6c5c-9a3440cba5ca")
	fmt.Println("Session的市信息市", sess)
}
