//闭包（closures）的形式之一

package main

import (
	"fmt"
)

//形式一
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//形式二
func sub() func() int {
	var x int
	return func() int {
		x--
		return x * x
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2 * i),
		)
	}
	sub := sub()
	fmt.Println(sub())
}