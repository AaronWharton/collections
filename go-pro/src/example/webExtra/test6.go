package main

import (
	"fmt"
	"html/template"
	"os"
)

type Friend struct {
	FriendName string
}

type People struct {
	PeopleName string
	Emails     []string
	Friends    []*Friend
}

func main() {
	// 条件处理if-else
	// ⚠️：与下面的警号标志结合起来看
	tEmpty, _ := template.New("template test").Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n")
	//tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	_ = tEmpty.Execute(os.Stdout, nil)
	// 空 pipeline if demo:

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	_ = tWithValue.Execute(os.Stdout, nil)
	// 不为空的 pipeline if demo:  我有内容，我会输出.

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	_ = tIfElse.Execute(os.Stdout, nil)
	// if-else demo:  if部分

	// Must操作：
	// 检测模版是否正确（大括号是否匹配等语法错误）
	// ⚠️：可与正则的Compile、MustCompile结合起来看：不用Must返回一个template和一个error，使用Must则忽略产生的error（直接panic），两者只是应用场景不同。
	// regexp.Compile()返回正则结构体和一个error，
	// 而regexp.MustCompile()只返回一个正则结构体。
	fmt.Println("new 1 parsed ok.")
	template.Must(template.New("new 1").Parse("new 1 parse and /* comment */"))

	fmt.Println("new 2 parsed ok.")
	template.Must(template.New("new 2").Parse("new 2 parse and variable {{.FriendName}}"))

	// 会产生panic，因为Must解析时发现了语法错误
	fmt.Println("new 3 caused panic because of variable {{.FriendName} not be placed correctly.")
	//template.Must(template.New("new 3").Parse("new 3 parse and variable {{.FriendName}"))

	// 输出嵌套字段内容，使用类似golang的for-range语法
	f1 := Friend{"Aaron"}
	f2 := Friend{"Allen"}
	t := template.New("new template file created")
	t, _ = t.Parse(`Hello {{.PeopleName}}!
	{{range .Emails}}
		an email {{.}}
	{{end}}
	{{with .Friends}}
	{{range .}}
		my friend name is {{.FriendName}}
	{{end}}
	{{end}}`)
	p := People{"people", []string{"zhd@qq.com", "zhd@163.com"}, []*Friend{&f1, &f2}}
	_ = t.Execute(os.Stdout, p)
}

// Output:
//Hello people!
//
//an email zhd@qq.com
//
//an email zhd@163.com
//
//
//
//my friend name is Aaron
//
//my friend name is Allen
//
//
