package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main(){
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("conn, err=",err)
		return
	}

	defer conn.Close() // 记住关闭链接

	// 通过go 向redis写入数据 string [key-value] 批量
	_,err = conn.Do("HMSet","user02","name","john","age",19)
	if err != nil{
		fmt.Println("HMSet err=",err)
		return
	}
	
	r,err := redis.Strings(conn.Do("HMGet","user02","name","age"))
	if err != nil{
		fmt.Println("HMGet err= ", err)
		return
	}

	// 遍历取
	// fmt.Printf("r=%v\n",r)
	for  i ,v := range r {
		fmt.Printf("r [%d]=%v\n",i,v)
	}
	
}