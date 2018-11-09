//defer的另一个实例

package main

import "fmt"

func main() {
	fmt.Println(test(1))
}

//这里输出1， 1。因为返回值是一个int，直接把function variable i 赋值给返回值（后续对i的修改不会影响返回值）。
//第一个输出是fmt的打印，然后执行return时调用defer，i++，然后返回返回值。
func test(i int) int {
	defer func() {
		i++
	}()
	fmt.Println(i)
	return i
}