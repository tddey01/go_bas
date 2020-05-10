package dao

import (
	"go_bas/go_web/day05/bookstore0612/model"
	"go_bas/go_web/day05/bookstore0612/utils"
	"net/http"
)

//AddSession  向数据中添加Session
func AddSession(sess *model.Session) error {
	// 写sql语句
	sqlStr := "INSERT INTO  sessions VALUES (?,?,?)"
	// 执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	// 写sql语句
	sqlStr := "DELETE  FROM  sessions WHERE session_id = ?"
	// 执行sql
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionById  根据session的ID值获取数据中查询Session
func GetSessionById(sessID string) (*model.Session, error) {
	// 写SQL语句
	sqlStr := "SELECT session_id ,username,user_id FROM  sessions WHERE session_id = ?"
	// 预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	// 执行
	row := inStmt.QueryRow(sessID)
	// 创建session
	sess := &model.Session{}
	//扫描数据库中的字段值为Session的字段赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	// 返回
	return sess, nil
}

// IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func IsLogin(r *http.Request) (bool, string) {
	// 根据cookie的name 获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		// 获取cookie的value
		cookievalue := cookie.Value
		// 根据cookievalue取数据中查询与之键对应的Session
		session, _ := GetSessionById(cookievalue)
		if session.UserID > 0 {
			//已经登录
			return true, session.UserName
		}
	}
	//没有登录
	return false, ""
}
