package main

import (
	"fmt"
	"os"
)

// 三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := If(a > b, a, b).(int)
	fmt.Println(result)

	// 以下代码会形成资源泄漏，因为在for循环里面使用defer，
	// 只会在for循环结束后调用最后一个defer语句而已。
	for i := 0; i < 5; i++ {
		f, err := os.Open("/xxx/xx/x.xx")
		if err != nil {
			// ignore error
		}
		// do something

		defer f.Close()
	}

	// 解决办法,在for循环里面使用func(){}()包装defer
	// 使得每次循环都在不同的函数域里面，从而每次循环都能调用defer释放资源。
	for i := 0; i < 5; i++ {
		func(){
			f,err := os.Open("xxx/xx/x.xx")
			if err != nil {
				// ignore error
			}
			// do something

			defer f.Close()
		}()
	}
}
