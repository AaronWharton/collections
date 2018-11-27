package main

import (
	"context"
	"fmt"
)

//func main() {
//	ch := func() <-chan int {
//		ch := make(chan int)
//		go func() {
//			for i := 0; ; i++ {
//				ch <- i
//			}
//		}()
//		return ch
//	}()
//
//	for v := range ch {
//		fmt.Println(v)
//		if v == 5 {
//			// 跳出循环后，后台goroutine无法退出，造成goroutine无法被回收
//			break
//		}
//	}
//}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:

				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}
