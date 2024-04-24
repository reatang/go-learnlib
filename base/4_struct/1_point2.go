package main

import "fmt"

// 结构体指针的特点2

//
// 1、指针接收器，用指针操作符取值
// 2、值接收器，创建了新值
//

type pB struct {
	a string
	b string
}

// New 此处使用了老变量创建了一个新的变量
func (p *pB) New(a string) *pB {
	_p := *p
	_p.a = a

	return &_p
}

// New2 这里和上面的功能一样，神奇的机制
func (p pB) New2(a string) *pB {
	p.a = a

	return &p
}

func main() {
	p1 := &pB{a: "a", b: "bb"}

	p2 := p1.New("b")
	p3 := p1.New2("c")

	fmt.Printf("%p, %+v\n", p1, p1)
	fmt.Printf("%p, %+v\n", p2, p2)
	fmt.Printf("%p, %+v\n", p3, p3)
}
