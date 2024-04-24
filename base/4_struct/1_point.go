package main

import "fmt"

// 结构体指针的特点

//
// 1、接收器内可以更换对象
//

type pA struct {
	a string

	pi *pI
}

type pI struct {
	b string
}

func (p *pA) replace(repl pA) {
	// p指针指向的内容可以更换
	*p = repl
}

func runP1() {
	p1 := pA{a: "ORIGIN"}
	fmt.Printf("%#v, %p\n", p1, &p1)

	// 此处可看出，值只是内容是替换了
	p2 := pA{a: "REPLACE"}
	p1.replace(p2)
	fmt.Printf("%#v, %p, %#v, %p\n", p1, &p1, p2, &p2)

	// 从此处看来，该操作没有申请新的内存空间，确实只是内容替换，属于浅拷贝
	pi := &pI{"PI"}
	p3 := pA{a: "P3P3P3", pi: pi}
	fmt.Printf("%#v, %p\n", p3, &p3)
	pi2 := &pI{"PI2"}
	p4 := pA{a: "P4", pi: pi2}
	p3.replace(p4)
	p4.a = "NEW P4" // P4不会影响 P3
	fmt.Printf("%#v, %p, %#v, %p\n", p3, &p3, p4, &p4)
	pi.b = "NEW PI"   // pi 已无关
	pi2.b = "NEW PI2" // 修改一个值，两个变量指向的内容都会改变
	fmt.Printf("%#v, %#v\n", p3.pi, p4.pi)

	fmt.Println("-----------")

	// 上面都是值操作，看看指针操作
	// 指针方法接收器，是对原值的新指针，而不是指针的指针
	prtp1 := &pA{a: "PRT P1"}
	fmt.Printf("%#v, %p, %p\n", prtp1, prtp1, &prtp1)
	prtp2 := pA{a: "PRT P2"}
	prtp1.replace(prtp2)
	fmt.Printf("%#v, %p, %p, %#v, %p\n", prtp1, prtp1, &prtp1, prtp2, &prtp2)
}

func main() {
	runP1()
}
