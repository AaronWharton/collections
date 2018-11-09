# Others

## 关于 slice 的一些思考

slice 三要素：指针（指向它代表的底层数组）、长度（声明时占有数据的长度）、容量（该 slice 所能容纳的数据量）。
slice 在重组的时候不能超过它的容量：
```go
package main

import "fmt"

func main() {
	a := [5]int{5, 5, 5, 5, 5}
	s1 := a[:]
	s2 := s1[:9]
	fmt.Println(len(s2))
}

// Output:
// panic: runtime error: slice bounds out of range

```
但是 slice 可以在 append 时超过它的容量（若超过容量，则直接扩大容量为当前的 2 倍）：
```go
package main

import "fmt"

func main() {
	a := [5]int{5, 5, 5, 5, 5}
	s1 := a[:]
	s2 := append(s1, 1, 2, 3, 4)
	fmt.Println(len(s2), cap(s2))
}

// Output:
// 9 10
```
