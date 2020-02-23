package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:7379")
	if err != nil {
		fmt.Println("Conn.Dial err=", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("HSet", "","name", "joyy", 10)
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	time.Sleep(time.Second * 11)
	r, err := redis.String("HGet", "name")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	fmt.Println(r)
}
