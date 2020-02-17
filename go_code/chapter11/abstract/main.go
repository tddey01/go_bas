package main

import "fmt"

// 定义一个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

//方法
// 1、 存款
func (account *Account) Desposite(money float64, pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.Pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入金额不正确")
		return
	}
	account.Balance += money
	fmt.Printf("存款成功 你当前账户余额=%v\n", account.Balance)
}

// 2、 取款
func (account *Account) WithDraw(money float64, pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.Pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 || money >= account.Balance {
		fmt.Println("你输入金额不正确")
		return
	}
	account.Balance -= money
	fmt.Printf("取款成功 你当前账户余额=%.2f\n", account.Balance)
}

// 查询余额
func (account *Account) Query(pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.Pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	fmt.Printf("你的账号为=%v 余额为=%v\n", account.AccountNo, account.Balance)
}

func main() {
	//测试
	account := Account{
		AccountNo: "gs11",
		Pwd:       "666",
		Balance:   100.00,
	}
	// for {
	//
	account.Query("666")
	account.Desposite(200, "666")
	account.WithDraw(299.99, "666")
	// }
}
