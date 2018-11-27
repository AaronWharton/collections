package test

import (
	"testing"
)

// 测试用例
// ⚠️：*_test.go文件执行时会检索当前目录下所有文件，
// 编写测试代码的时候，最好将他们放在一个单独的包内
// （例如：当前*.go文件都使用了main函数导致test时报错
// redeclare main function，当然正式项目不会发生多个
// 文件同时有main函数的情况...）。

func TestDivision(t *testing.T) {
	if result, err := Division(6, 2); result != 3 || err != nil {
		t.Error("Don't pass.")
	} else {
		t.Log("Passed.")
	}
}

// 压力测试
func BenchmarkDivision(b *testing.B) {
	// 下列两个函数是为了消除其他操作消耗时间影响压力测试的准确性
	// 停止计时
	//b.StopTimer()
	// 开始计时
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
