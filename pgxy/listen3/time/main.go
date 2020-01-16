package main

import (
	"fmt"
	"time"
)

func testTime() {
	new := time.Now()
	fmt.Println(new)
	fmt.Printf("current  time:%v\n", new)
	year := new.Year()
	month := new.Month()
	day := new.Day()
	hour := new.Hour()
	minute := new.Minute()
	send := new.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send) // 格式化
	//02d 不够两位 补全两位
	timestamp := new.Unix()

	fmt.Printf("timestamp %v\n", timestamp) // 时间戳
}

func testTimestamp(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()
	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
	// 时间戳格式化输入出
}

func processTesk() {
	fmt.Printf("do tesk\n")
}

func testtficker() {
	ticker := time.Tick(time.Second)
	// ticker := time.Tick(10*time.Second) // 每10秒一次
	for i := range ticker {
		fmt.Printf("ticker %v\n", i)
		processTesk()
	}
	// 每秒一次定时器
}

func testConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)   //纳秒
	fmt.Printf("micro second:%d\n", time.Microsecond) //微妙
	fmt.Printf("mili second:%d\n", time.Millisecond)  //毫秒
	fmt.Printf("second:%d\n", time.Second)            //秒
}

func testFormat() {
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:04:05\n")
	fmt.Printf("time:%s", timeStr)
	// 时间格式化输出  用于内部方法使用已经函数使用
}

func testFormat2() {
	now := time.Now()
	timeStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("time:%s", timeStr)
	// 时间格式化输出  用于内部方法使用已经函数使用
}

func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code cost :%d us\n", cost)

}
func main() {
	//testTime()   // 时间格式 输出
	// timestamp := time.Now().Unix()
	// testTimestamp(timestamp)  // 时间戳 转换 格式输出
	//testtficker()  // 定时器
	// testConst() //   时间纳秒 微妙 毫秒 秒 对比
	// testFormat() //格式化输出
	//testFormat2()
	testCost()
}
