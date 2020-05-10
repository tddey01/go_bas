package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "books:1234qwer@tcp(localhost)/books")
	if err != nil {
		panic(err.Error())
	}
}
