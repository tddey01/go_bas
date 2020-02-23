package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis" // 引入redis
)

func main() {
	// 1  链接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close() // 记得关闭链接
	// 通过go 向redis写入数据 string [key-value]
	_, err = conn.Do("HSet", "user01", "name", "john")
	if err != nil {
		fmt.Println("HSet err=", err)
		return
	}
	_, err = conn.Do("HSet", "user01", "age", 18)
	if err != nil {
		fmt.Println("HSet err=", err)
		return
	}

	// 通过 go 向redis 读取数据  string[key-value]
	r1, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil {
		fmt.Println("HGet err = ", err)
		return
	}
	r2, err := redis.String(conn.Do("HGet", "user01", "age"))
	if err != nil {
		fmt.Println("HGet err = ", err)
		return
	}

	//输出redis获取都数据库
	fmt.Printf("操作ok r1=%v  r2=%v\n", r1, r2)
}
