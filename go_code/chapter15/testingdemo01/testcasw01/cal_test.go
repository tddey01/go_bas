package cal

import (
	"fmt"
	"testing"
) // 引入go  的testing框架包

// 编写要给测试用例， 去测试addUpper是否正确
func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("addUpper错误 返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
	}
	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确..")
}

func TestHello(t *testing.T) {

	fmt.Println("TestHello被调用.....")

}
