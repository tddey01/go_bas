package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 定义一个全局pool
var pool *redis.Pool

// 当启动程序是， 就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大连接数， 0 表示没有限制
		IdleTimeout: 100, // 表示最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接的代码， 链接那个ip和端口
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	// 先从pool 池里面取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "tom猫")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//取出
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err= ", err)
		return
	}
	fmt.Println("Conn.Do r=", r)

	// 如果我们要从 pool 取出链接， 一定要保证连接池是没有关闭的
	pool.Close()
	conn2 := pool.Get()

	_, err = conn2.Do("Set", "name", "tom猫1")
	if err != nil {
		fmt.Println("conn2.Do err=", err)
		return
	}
	//取出
	r, err = redis.String(conn2.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn2.Do err= ", err)
		return
	}
	fmt.Println("Conn2.Do r=", r)
}
