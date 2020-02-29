package main

import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //这个表示指向前一个结点
	next     *HeroNode //这个表示指向下一个结点点
}

// 3、给双向链表插入一个节点
// 编辑第一个种插入方法， 在单链表中最后加入 [简单]
func InsterHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//	思路
	//	 1、 先找到该链表的最后这个节点
	//	 2、 创建一个复制节点[跑龙套, 帮助]
	temp := head
	for {
		if temp.next == nil { // 表示找到最后
			break
		}
		temp = temp.next // 让temp不断指向下一个节点
	}
	//	 3 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//给链表插入一个结点
//编写第2种插入方法，根据no 的编号从小到大插入..【实用】
func InsterHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//	思路
	//	 1、 找到适当的节点
	//	2 创建一个复制节点[跑龙套， 帮忙]
	temp := head
	flag := true
	// 让插入的节点的no 和temp的下一个节点的no比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no >= newHeroNode.no {
			// 说明newHeroNode 就应插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明我们额链表中已经有这个no,就不然插入.
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起 已经存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

// 删除一个节点
func DelHeroNode(head *HeroNode, id int) {

	temp := head
	flag := false
	// 找到要删除的节点的no 和temp的下一个节点的no比较
	for {
		if temp.next == nil { // 说明到链表的最后
			break
		} else if temp.next.no == id {
			//说明我们找到了.
			flag = true
			break
		}
		temp = temp.next
	}

	if flag { // 如果找到了 就删除
		temp.next = temp.next.next
	} else {
		fmt.Println("sorry 要删除的id不存在")
	}

}

// 显示聊表的所有节点信息
//这里仍然使用单向的链表显示方式
func ListHeroNode(head *HeroNode) {
	//	1 创建一个复制节点[跑龙套]
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也，.....")
		return
	}

	// 2  遍历这个链表
	for {
		fmt.Printf("[%d , %s %s] -->>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func ListHeroNode2(head *HeroNode) {
	//	1 创建一个复制节点[跑龙套]
	temp := head
	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也，.....")
		return
	}
	//2. 让temp定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	// 2  遍历这个链表
	for {
		fmt.Printf("[%d , %s %s] -->>", temp.no, temp.name, temp.nickname)
		//判断是否链表后
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	//	 1 、 先创建一个头结点, 默认
	head := &HeroNode{}
	//	 2、 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	// 3、 加入
	InsterHeroNode(head, hero1)
	InsterHeroNode(head, hero2)
	InsterHeroNode(head, hero3)

	ListHeroNode(head)

	fmt.Println("逆序打印")
	ListHeroNode2(head)

}
