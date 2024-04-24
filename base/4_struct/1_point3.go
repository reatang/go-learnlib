package main

import (
	"fmt"
)

// 结构体指针的特点3

//
// 实参指针 和 形参指针 是指向同一个地址的两个指针
//

type pC struct {
	a string
}

type pD []int

func instance(p *pC) {
	fmt.Printf("内层：%p, %p, %+v\n", p, &p, p)

	// 这样是无法更换本体的
	//p = &pC{a: "aaa"}
	//fmt.Printf("内层：%p, %p, %+v\n", p, &p, p)

	// 这样才能更换本体
	*p = pC{a: "aaabbb"}
}

func i2(p2 *pD) {
	fmt.Printf("%p, %+v\n", p2, p2)
	*p2 = append(*p2, 1)
	fmt.Printf("%p, %+v\n", p2, p2)
}

func main() {

	// 结论：p1 和 instance形参 p 只是指向了同一个内存值的两个指针，修改p不会影响p1
	var p1 = &pC{a: "000"}
	fmt.Printf("外层：%p, %p, %+v\n", p1, &p1, p1)
	instance(p1)
	fmt.Printf("外层：%p, %p, %+v\n", p1, &p1, p1)

	fmt.Println("-----------------------")

	// 如果想在函数内修改外部变量，必须使用&传入地址
	p2 := make(pD, 8, 8)
	fmt.Printf("%p, %+v\n", p2, p2)
	i2(&p2)
	//fmt.Printf("%d, %d\n", len(p2), cap(p2))
	fmt.Printf("%p, %+v\n", p2, p2)
}
