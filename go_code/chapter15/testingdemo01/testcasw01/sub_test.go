package cal

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	// 调用
	res := getSub(15, 2)
	if res != 13 {
		t.Fatalf("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 13, res)
	}
	// 如果正确，输出日志
	t.Logf("AddUpper(10) 执行正确.....")
}
