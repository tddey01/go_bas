package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

// 定义一个全局pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,     // 最大 连接数
		MaxActive:   maxActive,   // 表示和数据库的最大连接数， 0 表示无限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 初始化链接代码， 链接到那个ip的redis服务器
			return redis.Dial("tcp", address)
		},
	}
}
