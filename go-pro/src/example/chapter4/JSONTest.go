//JSON 练习

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	// field 必须以大写字母开头，否则解析后不显示数据
	Name   string
	Sex    string
	Grade  int
	Course []string
}

var stuInfo = []Person{
	{Name: "wmj", Sex: "nan", Grade: 1,
		Course: []string{"高数", "线代", "大物", "英语"}},
	{Name: "zhd", Sex: "nan", Grade: 2,
		Course: []string{"近代史"}},
	{Name: "lsh", Sex: "nan", Grade: 3,
		Course: []string{"那你很棒哦"}},
}

func main() {
	//data, err := json.Marshal(stuInfo) 不格式化，直接显示
	data, err := json.MarshalIndent(stuInfo, "", "	") //格式化json数据并显示
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
