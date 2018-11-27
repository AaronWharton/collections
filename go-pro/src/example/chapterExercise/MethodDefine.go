//method用法、nil是合法的receiver value

package main

import (
	"fmt"
	"math"
	"net/url"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func main() {
	r := Rectangle{12, 3}
	c := Circle{5}
	fmt.Println(r.area())
	fmt.Println(r) //参数传递variable副本，不改变原本的值
	fmt.Println(c.area())
	a := Rectangle{24, 30}
	a.change(2) //参数传递variable指针，改变variable值
	fmt.Println(a)

	/* 当m为nil时，它可进行Get操作（这不会添加/修改/删除数据）但是不能进行Add操作（会添加数据，map中的元素对caller可见，但是caller在这里为nil，所以错！） */
	m := url.Values{"item": {"en"}}
	m = nil
	fmt.Println(m.Get("item"))
	//m.Add("lang", "3")	//panic: assignment to entry in nil map
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (p *Rectangle) change(double float64) {
	p.height *= double
	p.width *= double
}
