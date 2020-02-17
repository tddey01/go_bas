package utils

import "fmt"

type FamliyAccount struct {
	// 声明必要字段
	// 生命一个字段 保存接收用户输入的选项
	key string
	// 声明一个字段 控制是否退出for循环
	loop bool
	//定义账户的余额[]
	balance float64
	// 每次收支的金额
	money float64
	//每次的收支说明
	note string
	// 定义个字段  记录是否有收支行为
	flag bool
	// 收支的详情使用字符串来记录
	// 当有收支时， 只需要对details 进行拼接处理
	details string // 空格符
}
// 编写一个工厂模式构造方法 返回一个FamilyAccount实例
func NewFamilyAccount() *FamliyAccount {
	return &FamliyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "收支\t账户金额\t收支金额\t说    明",
	}

}

// 将显示明细分装到到一个方法里面
func (this *FamliyAccount) showDetails() {
	fmt.Println("\n----------------------------当前收支明细记录-----------------------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细... 来一笔吧!")
	}
}

// 登记收入的封装到一个方法
func (this *FamliyAccount) income() {
	fmt.Println("本地收入金额")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本地收入说明:")
	fmt.Scanln(&this.note)
	// 将这个收入情况，拼接到details变量里面
	//收入    11000           1000            有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v \t  %v", this.balance, this.money, this.note)
	this.flag = true
}

// 将登记支出封装到一个方法
func (this *FamliyAccount) pay() {
	fmt.Println("本次支付的金额")
	fmt.Scanln(&this.money)
	//这里需要做一个必要判断
	if this.money > this.balance {
		fmt.Println("账户余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v \t  %v", this.balance, this.money, this.note)
	this.flag = true
}

// 将推出系统的封装到一个方法
func (this *FamliyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误 请重新输入？ y/n")
	}
	if choice == "y" {
		this.loop = false
	}
}

// 给该结构体绑定相应方法
//显示主菜单
func (this *FamliyAccount) MainMenu() {
	for {
		fmt.Println("--------------------------记账软件--------------------------")
		fmt.Println("------------------------1、收支明细--------------------------")
		fmt.Println("------------------------2、登记收入--------------------------")
		fmt.Println("------------------------3、登记支出--------------------------")
		fmt.Println("------------------------4、退出软件--------------------------")
		fmt.Println("请选择（1-4）:")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}
