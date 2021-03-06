package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串的操作
func main() {
	fmt.Println(strings.Contains("banana", "a")) // true
	fmt.Println(strings.Contains("banana", "s")) // false
	fmt.Println(strings.Contains("banana", ""))  // true
	fmt.Println(strings.Contains("", ""))        // true

	s := []string{"apple", "pear", "banana"}
	// 函数原型：func Join(a []string, sep string) string
	fmt.Println(strings.Join(s, ", ")) // Output: apple, pear, banana

	// 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        // ["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) // ["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         // [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))            // [""]

	// func Trim(s string, cutset string) string
	// 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]

	// func Fields(s string) []string
	// 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // Fields are: ["foo" "bar" "baz"]

	// ⚠️其他函数：

	// 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	// func Index(s, substr string) int

	// 重复s字符串count次，最后返回重复的字符串
	// func Repeat(s string, count int) string

	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	// func Replace(s, old, new string, n int) string

	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str)) // 4567false"abcdefg"'单'

	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e) // false 123.23 1234 12345 1023
}
