package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"runtime"
)

type Person struct {
	PersonName string
	email      string
}

func main() {
	// 调用导出字段会输出该字段的数据，
	// 当调用未导出字段（email）或者不存在的字段（Email）
	// 输出就为空而不报错（go version=go1.11.2）
	// 有的说未导出字段应该是报错，我猜是版本不同的原因
	t := template.New("new template")
	t, _ = t.Parse("Hello {{.PersonName}}!{{.Email}}!!{{.email}}!!!")
	s := Person{"Aaron", "zhd@qq.com"}
	if err := t.Execute(os.Stdout, s); err != nil {	// Output:Hello Aaron!
		log.Println(err)
	}
	// 检查当前使用的golang版本
	fmt.Println(runtime.Version())
}
