package model

import "fmt"

// 定义一个结构体 Account
type account struct {
	accountNo string
	pwd       string
	balance   float64
}

// 工厂模式一个函数
func NewAccount(accountNo string, pwd string, balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账户的长度不对")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码的长度不对")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额不能小于20")
		return nil
	}

	return &account{
		accountNo: accountNo,
		pwd:       pwd,
		balance:   balance,
	}
}

//方法
// 存款
func (account *account) Desposite(money float64, pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入金额不正确")
		return
	}
	account.balance += money
	fmt.Printf("存款成功 你当前账户余额=%v\n", account.balance)
}

// 2、 取款
func (account *account) WithDraw(money float64, pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	// 看看存款金额是否正确
	if money <= 0 || money >= account.balance {
		fmt.Println("你输入金额不正确")
		return
	}
	account.balance -= money
	fmt.Printf("取款成功 你当前账户余额=%.2f\n", account.balance)
}

// 查询余额
func (account *account) Query(pwd string) {
	// 看看输入密码是否福正确
	if pwd != account.pwd {
		fmt.Println("你输入密码不正确，请重新输入密码")
		return
	}

	fmt.Printf("你的账号为=%v 余额为=%v\n", account.accountNo, account.balance)
}
