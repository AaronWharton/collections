//可变参数函数 练习

package main

import (
	"fmt"
)

func main() {
	val := []int{25, -86, 45, -1, -12, -37, 95} //若输入为空则结果均为0
	fmt.Println(max(val...))                    //这里要把array转换为单个数据发送给max()，否则会发生类型异常
	fmt.Println(min(val...))                    //同上

	fmt.Println(double(4)) //defer实例
}

func max(values ...int) int {
	var max int
	for _, val := range values { //将输入缓冲区的第一个数据附给max
		max = val
		break
	}
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func min(values ...int) int {
	var min int
	for _, val := range values { //将输入缓冲区的第一个数据附给min
		min = val
		break
	}
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

/*关于 defer 详见 README defer*/
func double(x int) (result int) {
	defer func() {
		x += 4 //这里改变了x的值，但是不会影响上面的赋值，只会影响下面一行的输出（相当于只是取得了x的副本）
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	return x + x
}
