package controller

import (
	"go_bas/go_web/day03/bookstore0612/dao"
	"go_bas/go_web/day03/bookstore0612/model"
	"html/template"
	"net/http"

	"strconv"
)

// GetBooks 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// 调用bookdao中获取所有图书的函数
	books, _ := dao.GetBook()
	// 解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}

//addBook 添加图书
func AddBook(w http.ResponseWriter, r *http.Request) {
	//	获取图书信息
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// 将价格 销和库存进行转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)

	// 创建Book
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "static/img/default.jpg",
	}
	// 调用bookdao中的添加图书的函数
	dao.AddBook(book)
	//调用GetBooks处理器函数再次查询一次数据库
	GetBooks(w, r)

}

//DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//获取要删除的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中删除图书的函数
	dao.DeleteBook(bookID)
	//调用GetBooks处理器函数再次查询一次数据库
	GetBooks(w, r)
}

//  toUpdateBookPage 更新图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取要更新图书的id
	bookId := r.FormValue("bookId")
	// 调用bookdao中的话获取图书的函数
	book, _ := dao.GetBookByID(bookId)
	// 解析模板
	t := template.Must(template.ParseFiles("views/pages/manager/book_modify.html"))
	//执行
	t.Execute(w, book)
}

// UpdateBook 根据图书id更新图书信息
func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	//	获取图书信息
	bookid := r.PostFormValue("Id")
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	// 将价格 销和库存进行转换
	fbookid, _ := strconv.ParseInt(bookid, 10, 0)
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)


	// 创建Book
	book := &model.Book{
		Id:      int(fbookid),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "static/img/default.jpg",
	}
	// 调用bookdao中的添加图书的函数
	dao.UpdateBook(book)
	//调用GetBooks处理器函数再次查询一次数据库
	GetBooks(w, r)
}
