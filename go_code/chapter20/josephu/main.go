package main

import (
	"fmt"
)

// 小孩结构体
type Boy struct {
	No   int  // 编号
	Next *Boy // 指向下一个小孩的指针[默认是nil]
}

// 编写一个函数， 构成单向的环形链表
// num: 表示小孩的个数
// *Boy 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	// 构建第一个小孩
	first := &Boy{}   // 空指针
	curlBoy := &Boy{} // 空指针

	// 判断
	if num < 1 {
		fmt.Println("num的值不对")
		return first
	}
	// 开始循环的构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		//分析构成循环链表 需要一个辅助指针【帮忙的 ，】
		//1 因为第一个小孩比较特殊
		if i == 1 { // 第一个小孩
			first = boy
			curlBoy = boy
			curlBoy.Next = first // 形参一个循环指针
		} else {
			curlBoy.Next = boy
			curlBoy = boy
			curlBoy.Next = first // 构成环形链表
		}
	}
	return first
}

//显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {
	// 处理一下 环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}
	// 创建一个指针， 帮助遍历。
	curBoy := first
	for {
		//
		fmt.Printf("小孩的编号=%d -->\n", curBoy.No)
		// 退出条件  curBoy.Next == first
		if curBoy.Next == first {
			break
		}
		//curBoy移动到下一个
		curBoy = curBoy.Next
	}
}

/*
设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）
的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，
数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
*/

//分析思路
//1. 编写一个函数，PlayGame(first *Boy, startNo int, countNum int)
//2. 最后我们使用一个算法，按照要求，在环形链表中留下最后一个人
func PlayGame(first *Boy, starNo int, countNum int) {
	// 1 空的连接我们单独处理下
	if first.Next == nil {
		fmt.Println("空的连接 没有小孩")
		return
	}
	// 留一个 判断startNo <= 小孩的总数
	// 2 需要设置定义一个辅助的指针 帮助我们删除小孩
	tail := first
	// 3 下一步 让tail执行环形链表的最后一个小孩， 这个非常重要
	// 因为tail 在删除小孩时会用到
	for {
		if tail.Next == first { // 寿命tail到了最后小孩
			break
		}
		tail = tail.Next
	}
	// 4 让first移动到starNo [后面我们删除小孩，就以first为准]
	for i := 1; i <= starNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//	 5 、 开始数countNum ， 然后就删除first 指向的小孩
	for {
		// 开始数countNum - 1 次
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		// 删除first 执行的小孩
		first = first.Next
		tail.Next = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}
		fmt.Printf("小孩小孩编号为%d 出圈 \n", first.No)
	}
}

func main() {
	first := AddBoy(5)
	// 显示
	ShowBoy(first)
	PlayGame(first, 2, 1)
}
