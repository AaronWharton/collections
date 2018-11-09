// select 可以筛选数据，这个示例打印出奇数。注意他的原理：缓冲channel先send然后receive（channel空时receive会阻塞，满时send会阻塞）
// 但是 select 随机选取case。这是一个矛盾的问题（亟待解决）

package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	for i := 1; i < 11; i++ {
		select {
		case x:= <- ch:
			fmt.Println(x, "++++", i)
		case ch <- i:
			//fmt.Println("这里先触发......", i)
		}
	}
}
