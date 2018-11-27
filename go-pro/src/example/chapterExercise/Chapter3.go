//第三章练习题

package main

import "fmt"

const (
	i = 1 << iota //iota为0，初始值为1（0001）左移0位，i输出为1
	j = 3 << iota //iota递增为1，初始值为3（0011），左移1位，j输出为6
	k             //iota递增为2，初始值为3（0011），左移2位，k输出为12
	l             //iota递增为3，初始值为3（0011），左移3位，l输出为24
)

func main() {
	/* iota 的使用*/
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	/* range 的用法*/
	fmt.Println("---------------------------------------")
	nums := []int{2, 3, 4}
	for i, num := range nums {
		//i为索引，num为对应的值
		if num == 3 {
			fmt.Println("item:", i) //item: 1
		}
	}
}
