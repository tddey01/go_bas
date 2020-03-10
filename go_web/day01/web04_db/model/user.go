package model

import (
	"fmt"
	"go_bas/go_web/day01/web04_db/utils"
)

// User 结构体
type User struct {
	Id       int
	Username string
	Passord  string
	Email    string
}

// 添加 ADDUser 添加User的方法 一
func (user *User) AddUser() error {
	// 写 SQL 语句
	sqlStr := "INSERT INTO users(username,passord,email) VALUES (?,?,?)"
	// 预编译
	InStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常", err)
		return err
	}
	// 执行sql语句
	_, err2 := InStmt.Exec("admin", "123456", "admin@gu.com")
	if err2 != nil {
		fmt.Println("执行SQL异常", err)
		return err2
	}
	return nil
}

// 添加 ADDUser 添加User的方法 二
func (user *User) AddUser2() error {
	// 写 SQL 语句
	sqlStr := "INSERT INTO users(username,passord,email) VALUES (?,?,?)"
	// 执行sql语句
	_, err := utils.Db.Exec(sqlStr, "admin2", "123456", "admin2@gu.com")
	if err != nil {
		fmt.Println("添加User的方法 二 执行SQL异常", err)
		return err
	}
	return nil
}
