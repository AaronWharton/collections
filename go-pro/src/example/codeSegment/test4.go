package main

import (
	"fmt"
	"sync"
)

// 不正确的写法，多个goroutine访问同一个数据（这里为count）
// 会导致数据丢失(这里不能保证输出1000)，竞争访问应该对访问的对象加上锁。
//
//import (
//	"fmt"
//	"sync"
//)
//
//var count int
//
//func main() {
//	var wg sync.WaitGroup
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			count++
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println(count)
//}

// 加锁后保证能输出1000
//
var count int

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
