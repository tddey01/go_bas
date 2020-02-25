package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 我们再发服务器启动后， 就初始化一个userDao实例
// 把它做成全局的变量， 在需要和redsi操作时， 就直接使用即可
var (
	MyUserDao *UserDao
)

// 定义一个UserDao 结构体
// 完成对User 结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式 创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//思考一下在UserDao 应该提供哪些方法给我们
// 1 、 根据用户id  返回 一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定id 去redis 里面查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 错误
		if err == redis.ErrNil { // 表示在users  哈希中没有找到对应id
			err = ERROR_USER_NOTEXSTS // 返回自定义错误信息
		}
		return
	}
	user = &User{}
	// 我们需要把res 反序列化成User实例
	//err = json.Unmarshal([]byte(res), user)
	fmt.Println(res)
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), user) err=", err)
		return
	}
	return
}

// 完成登录校验 login
// 1 Login 完成对用户的校验
// 2 如果用户的id和pwd都正确  怎返回一个user实例
// 3 如果用户的id和pwd  有错误， 怎返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	// 先从UserDao 的连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	// 调用 this.getUserById(conn,userId)
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	// 这时证明用户我们是获取到了 ， 但是他的密码校验
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
