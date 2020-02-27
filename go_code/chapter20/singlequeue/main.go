package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组=> 模拟队列
	front   int    // 表示指向队列首
	rear    int    // 表示指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQuQueue(val int) (err error) {
	// 先判断队列是否已满
	if this.rear == this.maxSize-1 { // 重要的提示； rear 是队列尾(含最后的元素)
		return errors.New("Queue Full 满了")
	}

	this.rear++ // rear 后移
	this.array[this.rear] = val
	return
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.rear == this.front { // 对列空了
		return -1, errors.New("Queue empty ")
	}
	this.front++
	val = this.array[this.front]
	return val, err

}

//  显示队列 找到对首 ， 然后遍历到对尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	// this.front 不包含对首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t\n", i, this.array[i])
	}

}

// 编写一个主函数 测试
func main() {
	//	先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}

	// 添加数据

	var key string
	var val int
	for {
		fmt.Println("1 输入 add 表示添加数据到队列")
		fmt.Println("2 输入 get 表示队列中获取数据")
		fmt.Println("3 输入 show 表示显示队列")
		fmt.Println("3 输入 exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加到队列的数")
			fmt.Scanln(&val)
			err := queue.AddQuQueue(val)
			if err != nil {

				fmt.Println("加入队列失败", err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("从队列中取出了一个数=%d\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误, 请重新输入")

		}

	}
}
