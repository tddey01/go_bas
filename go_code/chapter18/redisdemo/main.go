package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis" // 引入redis
)

func main() {
	// 通过go 向redis  写入数据和读取数据
	// 1 链接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() // 记住关闭
	// 通过go向redis写入数据string[key-value]
	_, err = conn.Do("set", "name", "tomjerrym喵喵")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	// 通过go 向redis读取数据 string [key-value]
	// r, err := conn.Do("Get", "name")
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("Get err=", err)
		return
	}
	// 因为返回r‘是interface{}
	// 因为name 对应的执行时string，因此我们需要转换
	// nameString := r.

	fmt.Println("conn succ...", r)
}
