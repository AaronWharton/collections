//第四章练习题

package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {

	/* 数组的特点 */
	fmt.Println("---------------------------------------")
	q := []int{1, 2, 3}
	fmt.Printf("%T\n", q) //[]int
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p) //[3]int

	/* slice 移除元素 */
	fmt.Println("---------------------------------------")
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	/* 在 map 中查找不存在的键 */
	fmt.Println("---------------------------------------")
	ages := make(map[string]int)
	ages["Alice"] = 314
	ages["Aaron"] = 215
	fmt.Println(ages["Bob"]) //若查找不存在的键不会报错，返回""或者0
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	/* 不能向未初始化（默认初始化）的 map 中添加数据，详见 Documents 7 */
	fmt.Println("---------------------------------------")
	//var aa map[int]string
	bb := make(map[int]string)
	//aa[1] = "Alice" //未初始化，不能向其添加数据，否则报错。
	bb[1] = "Aaron" //使用了make()为它开辟内存空间，可以向它添加数据
	//fmt.Println(aa) //map[]
	fmt.Println(bb) //map[1:Aaron]

	/* struct 赋值，若未给指定 field 赋值，则需要将每个 field 的值都初始化 */
	fmt.Println("---------------------------------------")
	//r := Point{1, 2} //没有给指定 field 初始化，则需要将全部 field 初始化
	r := Point{X : 1} //指定给 x 初始化，则可以不用给 y 初始化（此时 y 的值为默认0）
	fmt.Println(r.X, r.Y)

	/* struct 的比较 */
	fmt.Println("---------------------------------------")
	pp := Point{1, 2}
	qq := Point{2, 1}
	fmt.Println(pp == qq) //只有在同一个 struct 下才可以进行比较，当 struct 下的对应的 field 值都相等时为true，否则为false
}

func remove(slice []int, i int) []int {
	//copy(slice[i:], slice[i+1:]) //有序
	slice[i] = slice[len(slice) - 1] //无序
	return slice[:len(slice) - 1]
}