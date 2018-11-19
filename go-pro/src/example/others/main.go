package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

type struct1 struct {
	field1 string
	field2 int
}

type Stringer1 int

func main() {

	// 字符串
	//
	// 如何修改字符串中的一个字符？如何获取一个字符串的字符数？
	//
	// 注：str1包含非ASCII字符，而下面的替换字符只能替换ASCII字符
	// utf8.RuneCountInString(str1)统计字符个数而不是字节数，因此输出5
	str1 := "😄ello"
	c := []byte(str1)
	c[0] = 'c'
	c1 := string(c)
	fmt.Println(c1, utf8.RuneCountInString(str1)) // Output:c���ello 5

	// 映射
	//
	// 如何对map初始化、删除map键以及判断某个键是否存在？
	str2 := make(map[string]int)
	str2 = map[string]int{
		"str":  1,
		"str1": 2,
		"str2": 3,
		"str3": 4,
	}
	delete(str2, "str1")
	k, isExist := str2["str1"]
	if !isExist {
		fmt.Println("key not exist!") // Output:key not exist!
	} else {
		fmt.Println(k)
	}

	// 结构体
	//
	// 如何创建和初始化struct？
	ms1 := new(struct1)                       // 注：以此方式创建struct时会给成员附初值，string=""， int=0 etc.
	fmt.Println(ms1.field1, "\n", ms1.field2) // Output: 2（注：field1为""值）
	ms2 := NewStruct1("This is a new struct", 1)
	fmt.Println(ms2.field1, ms2.field2) // Output:This is a new struct 1

	// 接口
	//
	// type assertion
	// 如何检测一个值是否实现了接口Stringer？
	var v interface{}
	if v, ok := v.(Stringer1); ok {
		fmt.Println("implements Stringer1:", v)
	} else {
		fmt.Println("not implement Stringer1") // Output:not implement Stringer1
	}

	// 类型判断
	var1, var2, var3, var4 := 1, "var2", 3.14, false
	classifier(var1, var2, var3, var4)
}

// struct的构建函数，通常使用该函数对struct变量初始化，通常struct都需要占用大量内存空间，因此使用*struct
func NewStruct1(s string, i int) *struct1 {
	return &struct1{s, i}
}

func classifier(items ...interface{}) {
	for _, typ := range items {
		switch typ.(type) {
		case int:
			fmt.Println(typ, "is int type.")
		case string:
			fmt.Println(typ, "is string type.")
		case bool:
			fmt.Println(typ, "is bool type.")
		default:
			log.Println("unknown type of variable:", typ)
		}
	}
}
