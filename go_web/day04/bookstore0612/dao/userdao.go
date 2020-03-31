package dao

import (
	"go_bas/go_web/day04/bookstore0612/model"
	"go_bas/go_web/day04/bookstore0612/utils"
)

// CheckUserNameAndPassword 根据用户和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	// 写SQL语句
	sqlStr := "SELECT id,username,PASSWORD,Email FROM users WHERE username = ? AND PASSWORD = ?"
	//	执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// CheckUserName 根据用户和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User, error) {
	// 写SQL语句
	sqlStr := "SELECT id,username,PASSWORD,Email FROM users WHERE username = ?"
	//	执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, email string) error {
	// 写SQL语句
	sqlStr := "INSERT INTO  users( username, PASSWORD, Email)  VALUES (?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
