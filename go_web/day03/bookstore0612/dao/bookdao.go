package dao

import (
	"go_bas/go_web/day02/bookstore0612/utils"
	"go_bas/go_web/day03/bookstore0612/model"
)

// GetBook 获取所有数据库中所有的图书
func GetBook() ([]*model.Book, error) {
	// SQL语句
	sqlStr := "SELECT id,title,author,price,sales,stock,img_path FROM books.book"
	// 执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, nil
}

// AddBook  向数据库中添加一个本图书
func AddBook(b *model.Book) error {
	// 写SQL 语句
	sqlStr := "INSERT INTO book ( title, author, price, sales, stock, img_path)  VALUES  (?,?,?,?,?,?)"
	// 执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}