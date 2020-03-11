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

// GetUserById 根据用户的id从数据库中查询一条记录
func (user *User) GetUserById() (*User, error) {
	// 写SQL语句
	sqlStr := "SELECT  id,username,passord,email FROM users WHERE id = ?"
	// 执行
	row := utils.Db.QueryRow(sqlStr, user.Id)
	// 声明
	var (
		id       int
		username string
		passord  string
		email    string
	)
	err := row.Scan(&id, &username, &passord, &email)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:       id,
		Username: username,
		Passord:  passord,
		Email:    email,
	}
	return u, nil
}

// GetUsersAll 获取数据库所有的的记录
func (user *User) GetUserAll() ([]*User, error) {
	// 写SQL语句
	sqlStr := "SELECT  id,username,passord,email FROM users "
	// 执行
	row, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	// for 循环
	// 创建User切片
	var users []*User
	for row.Next() {
		var (
			id       int
			username string
			passord  string
			email    string
		)
		err := row.Scan(&id, &username, &passord, &email)
		if err != nil {
			return nil, err
		}
		u := &User{
			Id:       id,
			Username: username,
			Passord:  passord,
			Email:    email,
		}
		users = append(users, u)
	}
	return users, nil
}
