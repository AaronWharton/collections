package main

import (
	"fmt"
)

var msg string
var done1 bool
var done2 = make(chan bool)

func main() {
	// 传入空接口的切片时需要注意参数展开的问题
	s := []interface{}{1, 2, 3}
	fmt.Println(s)    // [1 2 3]
	fmt.Println(s...) // 1 2 3

	// recover结合panic必须在defer里面使用
	//defer func() {
	//	recover()
	//}()
	//panic(1)

	// 不同goroutine之间不满足顺序一致性内存模型，
	// 因此在不同goroutine之间最好通过channel传递数据
	//go setup1()
	//for !done1 {
	//}
	//fmt.Println(msg)
	go setup2()
	select {
	case <-done2:
		fmt.Println(msg)
	}
}

func setup2() {
	msg = "Hello world!"
	done2 <- true
}

func setup1() {
	msg = "Hello world"
	done1 = true
}
